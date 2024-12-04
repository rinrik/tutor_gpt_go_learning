package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	numberOfGorutines := 5
	ch := make(chan []int)

	for i := 0; i < numberOfGorutines; i++ {
		wg.Add(1)
		go handleArray(ch, &wg)
	}

	for {
		slice := generateIntSlice()
		if len(slice) <= 1 {
			break // Останавливаем цикл, если длина массива <= 1
		}
		ch <- slice
	}

	close(ch)
	wg.Wait()

}

func generateIntSlice() []int {
	var res []int
	rand.NewSource(time.Now().Unix())
	randomSliceLen := rand.Intn(6)

	for i := 0; i <= randomSliceLen; i++ {
		randomDigit := rand.Intn(100)
		res = append(res, randomDigit)
	}

	return res
}

func handleArray(ch chan []int, wg *sync.WaitGroup) {
	defer wg.Done()

	for arr := range ch { // Читаем из канала до его закрытия
		var sum int
		for _, v := range arr {
			sum += v
		}
		log.Printf("The sum of array is: %d", sum)
	}
}
