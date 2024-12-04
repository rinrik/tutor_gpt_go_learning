package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Type  string
	Value string
}

type WorkerChan struct {
	UpperWork chan string
	LowerWork chan string
}

func (t *Task) validateTaskType() (bool, error) {
	switch strings.ToLower(t.Type) {
	case "upper":
		return true, nil
	case "lower":
		return true, nil
	case "help":
		return false, fmt.Errorf("'upper' and 'lower' types are available")
	case "exit":
		os.Exit(0)
		return false, nil
	default:
		return false, fmt.Errorf("type help, to see available commands")
	}
}

func main() {
	var task Task
	query := make(chan Task)
	workers := WorkerChan{
		UpperWork: make(chan string),
		LowerWork: make(chan string),
	}

	go taskListener(query, workers)
	go taskWorker(workers)

	for {
		fmt.Println("Enter task type")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		task.Type = scanner.Text()

		_, err := task.validateTaskType()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Println("Enter task value")
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		task.Value = scanner.Text()
		query <- task
	}

}

func taskListener(ch chan Task, workCh WorkerChan) {
	for t := range ch {
		switch t.Type {
		case "upper":
			workCh.UpperWork <- t.Value
		case "lower":
			workCh.LowerWork <- t.Value
		}
	}
}

func taskWorker(work WorkerChan) {
	var textToProceed string
	select {
	case textToProceed = <-work.UpperWork:
		fmt.Printf("Job is done: %v\n", strings.ToUpper(textToProceed))
	case textToProceed = <-work.LowerWork:
		fmt.Printf("Job is done: %v\n", strings.ToLower(textToProceed))

	}
}
