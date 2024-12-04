package main

import "fmt"

func main()  {
	var arrLength int
	var input int

	fmt.Println("Введите кол-во цифр в масиве")
	_, err := fmt.Scanln(&arrLength)
	if err != nil {
		fmt.Println("Ошибка: введено неверное значение длины масива")
	}

	arr := make([]int,arrLength)

	for i:=0; i<arrLength; i++ {
		fmt.Printf("введите цифру на позиции: %d\n", i)
		fmt.Scan(&input)
		arr = append(arr, input)
	}

	bubbleSort(&arr)

	fmt.Printf("Минимальное значени в массиве - %d\n", arr[0])
	fmt.Printf("Максимально значени в массиве - %d\n", arr[len(arr)-1])
}


func bubbleSort(arr2 *[]int) {
	arr := *arr2
	n := len(arr)
	for i := 0; i < n-1; i++ {
	    for j := 0; j < n-i-1; j++ {
		if arr[j] > arr[j+1] {
		    arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	    }
	}
    }