package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const dateFormat = "2 Jan 2006"

type currentTime struct {
	Time string `json:"current-time"`
}

func main() {
	http.HandleFunc("/", timeHandlerFunc)
	http.HandleFunc("/delay", delayHandlerFunc)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server can't start: %v", err)
	}
}

func timeHandlerFunc(resp http.ResponseWriter, req *http.Request) {
	currentTime := currentTime{
		Time: time.Now().Format(dateFormat),
	}

	bytedCurrentTime, err := json.Marshal(currentTime)
	if err != nil {
		log.Fatalf("Can't encode JSON current time %v", err)
	}

	resp.Write(bytedCurrentTime)
}

func delayHandlerFunc(resp http.ResponseWriter, req *http.Request) {
	type errorResponse struct {
		Error string `json:"error"`
	}

	interruptedError := errorResponse{
		Error: "Request was interrupted",
	}

	// Создаем канал для таймера
	timer := time.NewTimer(5 * time.Second)
	ctx := req.Context()

	select {
	// Если клиент отменяет запрос
	case <-ctx.Done():
		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(http.StatusRequestTimeout)
		logError, _ := json.Marshal(interruptedError)
		resp.Write(logError)
		fmt.Println(interruptedError.Error)
		return

	// Если таймер истек
	case <-timer.C:
		resp.Header().Set("Content-Type", "text/plain")
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte("Delay completed!"))
		return
	}
}
