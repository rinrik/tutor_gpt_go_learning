package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Создаем слушающий сокет
	listener, err := net.Listen("tcp", ":8056")
	if err != nil {
		fmt.Println("Ошибка создания сервера:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Сервер слушает на порту 8080...")
	go client()

	for {
		// Принимаем входящее соединение
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка подключения:", err)
			continue
		}

		// Обрабатываем соединение в отдельной горутине
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Новое соединение от:", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	for {
		// Читаем данные от клиента
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения:", err)
			break
		}

		fmt.Printf("Получено сообщение: %s", message)

		// Отправляем ответ клиенту
		_, err = conn.Write([]byte(strings.ToUpper(message)))
		if err != nil {
			fmt.Println("Ошибка отправки:", err)
			break
		}
	}
}

func client() {
	// Устанавливаем соединение с сервером
	conn, err := net.Dial("tcp", "localhost:8056")
	if err != nil {
		fmt.Println("Ошибка подключения к серверу:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Подключено к серверу. Введите сообщения для отправки (для выхода наберите `exit`)")

	reader := bufio.NewReader(os.Stdin)
	for {
		// Читаем ввод от пользователя
		fmt.Print("Введите сообщение: ")
		message, _ := reader.ReadString('\n')

		if message == "exit\n" {
			fmt.Println("Завершение соединения.")
			break
		}

		// Отправляем сообщение серверу
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Ошибка отправки:", err)
			break
		}

		// Получаем ответ от сервера
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения ответа:", err)
			break
		}
		fmt.Printf("Ответ от сервера: %s", response)
	}
}
