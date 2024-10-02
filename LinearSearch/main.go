package main

import (
	"fmt"
	"time"
)

func main() {
	counter := 0

	number := 10000000

	array1 := GenerateSortedArray(100000000)
	
	start := time.Now()
	LinearSearchWithArray(array1, number, &counter)
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s - Counter: %d\n", elapsed, counter)
}

func GenerateSortedArray(size int) []int {
	array := make([]int, size)

	for i := 0; i < size; i++ {
		array[i] = i + 1
	}

	return array
}

func LinearSearchWithArray(input []int, searchNumber int, counter *int) bool {
	for _, value := range input {
		*counter++
		if value == searchNumber {
			fmt.Println("--------------------------SEARCH FINISHED-----------------------------")
			return true
		}
	}

	return false
}
