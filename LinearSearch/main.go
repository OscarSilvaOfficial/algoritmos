package main

import "fmt"

func LinearSearchWithArray (input [10]int) {
	fmt.Println("--------------------------------------------------------")

	for index := 0; index < 10; index++ {
		message := fmt.Sprintf("Index: %d, Number: %d", index, input[index])

		if index >= 0 && index <= 3 {
			fmt.Println(message)
		}
	}

	for index := 0; index < len(input); index++ {
		message := fmt.Sprintf("Index: %d, Number: %d", index, input[index])
		if index > 3 && index <= 6 {
			fmt.Println(message)
		}
	}

	for index, value := range  input {
		message := fmt.Sprintf("Index: %d, Number: %d", index, value)
		if index > 6 && index <= 10 {
			fmt.Println(message)
		}
	}

	fmt.Println("--------------------------------------------------------")
}

func main() {
	var array1 [10]int = [10]int{1, 2, 30, 12, 34, 95, 21, 22, 44, 12} 
	LinearSearchWithArray(array1)
}