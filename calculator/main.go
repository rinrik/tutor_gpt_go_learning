package main

import (
    "fmt"
)

func main() {
    var firstValue, secondValue int
    var operator string
    var result int

    fmt.Print("Enter first value: ")
    _, err := fmt.Scanln(&firstValue)
    if err != nil {
        fmt.Println("Invalid input. Please enter an integer.")
        return
    }

    fmt.Print("Choose one operator: + - * /: ")
    fmt.Scanln(&operator)

    fmt.Print("Enter second value: ")
    _, err = fmt.Scanln(&secondValue)
    if err != nil {
        fmt.Println("Invalid input. Please enter an integer.")
        return
    }

    switch operator {
    case "+":
        result = firstValue + secondValue
    case "-":
        result = firstValue - secondValue
    case "*":
        result = firstValue * secondValue
    case "/":
        if secondValue == 0 {
            fmt.Println("Error: Division by zero.")
            return
        }
        result = firstValue / secondValue
    default:
        fmt.Println("Invalid operator. Please choose one of +, -, *, /.")
        return
    }

    fmt.Printf("Result is: %d\n", result)
}
