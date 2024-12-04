package main

import (
	"fmt"
	"os"
	"sync"
)

// если запускаешь из терминала, то надо убрать workDir
func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var textCount int

	fileName := make(chan string)

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go checkFileContent(fileName, &wg, &textCount, &mu)
	}

	workDir := "task_16_file_processing_with_gorutines"
	filesFolder := "files"
	filesPath := fmt.Sprintf("./%s/%s", workDir, filesFolder)

	files, err := os.ReadDir(filesPath)
	if err != nil {
		fmt.Printf("ERROR: can't read files directory. err: %s\n", err)
		panic(err)
	}

	for _, file := range files {
		fileName <- fmt.Sprintf("%s/%s", filesPath, file.Name())
	}
	close(fileName)

	wg.Wait()

	fmt.Printf("Files with text 'lorem ipsum': %v", textCount)

}

func checkFileContent(fNames chan string, wg *sync.WaitGroup, counter *int, mu *sync.Mutex) {
	expectedText := "lorem ipsum"
	defer wg.Done()
	for fName := range fNames {
		fileText, err := os.ReadFile(fName)
		if err != nil {
			fmt.Printf("ERROR: can't read file. err: %v", err)
			panic(err)
		}

		if expectedText == string(fileText) {
			mu.Lock()
			*counter++
			mu.Unlock()
		}

	}
}
