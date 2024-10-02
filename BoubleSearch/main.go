package main

import (
	"fmt"
	"time"
)

func main() {
	counter := 0

	number := 100000000

	array1 := GenerateSortedArray(100000000)
	
	start := time.Now()
	data := BoubleSearchWithArray(array1, number, &counter)
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s - Counter: %d\n", elapsed, counter)
	fmt.Printf("Number: %d", data)
}

func BoubleSearchWithArray(array []int, searchNumber int, counter *int) int {
	*counter++

	if len(array) == 1 {
		return array[0]
	}

	left := array[:len(array) / 2]
	right := array[len(array) / 2:]

	if left[len(left)-1] > searchNumber {
		return BoubleSearchWithArray(left, searchNumber, counter)
	}
	
	return BoubleSearchWithArray(right, searchNumber, counter)
}

func GenerateSortedArray(size int) []int {
	array := make([]int, size)

	for i := 0; i < size; i++ {
		array[i] = i + 1
	}

	return array
}