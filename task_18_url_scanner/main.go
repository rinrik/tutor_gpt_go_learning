package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

func main() {
	ur := make(chan string)
	var urls []string
	var wg sync.WaitGroup

	fmt.Println("Enter url list with following rules \n - Url should be without https|http and :// \n - Enter URLs one by one \n - To start scanning type START")

	for {
		var str string
		_, err := fmt.Scan(&str)
		if err != nil {
			fmt.Errorf("Invalid entered URL %v", err)
			break
		}
		if str == "START" {
			fmt.Println("Scanner started")
			break
		}
		urls = append(urls, str)
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go scanner(ur, &wg)
	}

	for _, v := range urls {
		ur <- v
	}
	close(ur)

	wg.Wait()

}

func scanner(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	client := new(http.Client)
	for url := range ch {
		resp, err := client.Get(strings.Join([]string{"https://", url}, ""))
		if err != nil {
			fmt.Printf("ERROR: Can't open %v, because: %v \n", url, err)
			continue
		}

		if resp.StatusCode != 200 {
			fmt.Printf("Url - %v is failed with status %v \n", url, resp.StatusCode)
		} else {
			fmt.Printf("Url - %v is successful with status %v \n", url, resp.StatusCode)
		}
	}
}
