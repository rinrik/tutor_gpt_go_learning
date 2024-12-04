package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2/log"
)

func main() {
	urlCh := make(chan string)
	isDone := make(chan bool)

	fmt.Println("Enter url or several urls with space as divider")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	go checkUrlStatus(urlCh, isDone)

	urls := strings.Split(input, " ")
	for _, v := range urls {
		urlCh <- v
	}

	close(urlCh) // закрываем канал когда все URL переданы

	<-isDone                               // сука ждем заверщения канала, как же я намучался тут
	fmt.Println("Все проверки завершены.") // спокойно выходим

}

func checkUrlStatus(ch chan string, done chan bool) {
	client := new(http.Client)
	for url := range ch {
		resp, err := client.Get(strings.Join([]string{"https://", url}, ""))
		if err != nil {
			log.Errorf("Can't do request to: %v, because of %v", url, err)
			continue
		}
		status := strings.Join([]string{url, " status is ", resp.Status}, "")
		log.Info(status)
	}
	done <- true
}
