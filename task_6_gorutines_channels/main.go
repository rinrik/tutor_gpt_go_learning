package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan int)

	go func() {
		for i:=1; i<=5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()
	
	for msg := range ch {
		fmt.Println(msg)
	}


}