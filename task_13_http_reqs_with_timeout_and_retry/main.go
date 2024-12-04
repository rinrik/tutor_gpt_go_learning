package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Enter URL's to call, with space as separator")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()
	urls := strings.Split(input, " ")
	beatufyUrls(&urls)

	for _, v := range urls {
		getRequestWithRetries(v, 3) // Передаем количество попыток (например, 3)
	}
}

func beatufyUrls(arr *[]string) {
	dummyArr := *arr
	for i, v := range dummyArr {
		dummyArr[i] = strings.Join([]string{"https://", v}, "")
	}
}

func getRequestWithRetries(url string, maxRetries int) {
	var lastErr error
	for i := 0; i < maxRetries; i++ {
		// Выполняем запрос с таймаутом
		err := getRequest(url)
		if err == nil {
			return // Если запрос выполнен успешно, выходим из функции
		}

		// Если была ошибка, сохраняем ее и ждем перед следующей попыткой
		lastErr = err
		log.Printf("Retry %d/%d for URL: %s failed with error: %v", i+1, maxRetries, url, err)

		// Задержка перед следующей попыткой
		time.Sleep(500 * time.Millisecond)
	}

	// Если после всех попыток запрос не удался, выводим ошибку
	log.Printf("All retries failed for URL: %s, last error: %v", url, lastErr)
}

func getRequest(url string) error {
	// Устанавливаем таймаут на запрос
	timeOut := 3 * time.Second // Можно изменить таймаут
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	// Создаем HTTP-запрос с контекстом
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Отправляем запрос
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// Проверяем, истек ли таймаут
		if ctx.Err() == context.DeadlineExceeded {
			return fmt.Errorf("request timed out")
		}
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Если запрос выполнен успешно, выводим статус
	log.Printf("SUCCESS\n URL: %v\n STATUS:%v", url, resp.StatusCode)
	return nil
}
