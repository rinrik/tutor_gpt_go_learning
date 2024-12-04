package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    var input string
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Введите свой текст:")
    input, err := reader.ReadString('\n')
    if err != nil {
        panic(err)
    }

    err = os.WriteFile("text.txt", []byte(input),0666)
    if err != nil {
        panic(err)
    }

    writtenFile, err := os.ReadFile("text.txt")
    if err != nil {
        panic(err)
    }

    fmt.Println("Вот что вы записали в файл:")
    fmt.Println(string(writtenFile))
}
