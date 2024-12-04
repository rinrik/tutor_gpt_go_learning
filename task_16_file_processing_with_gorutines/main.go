package main

import (
	"fmt"
	"os"
)

func main() {
	var fileTexts [][]byte
	files, _ := os.ReadDir("./files")
	fmt.Println(files[0].Name())
	for _, v := range files {
		temp, _ := os.ReadFile("files/" + v.Name())
		fileTexts = append(fileTexts, temp)

	}

	for _, f := range fileTexts {
		fmt.Println(string(f))
	}
}
