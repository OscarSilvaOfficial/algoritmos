package main

import (
	"fmt"
	"time"
	"algoritmos/search"
	"algoritmos/utils"
)

func main() {
	counter := 0

	number := 1
	array1 := utils.GenerateSortedArray(511)
	
	// IMPORTANT: Binary Search
	start := time.Now()
	data := search.BoubleSearch(array1, number, &counter)
	elapsed := time.Since(start)

	fmt.Printf("Binary Search: Execution time: %s - Counter: %d\n", elapsed, counter)
	fmt.Printf("Number found: %d", data)
	fmt.Printf("\n\n")

	// IMPORTANT: Linear Search
	counter = 0

	start = time.Now()
	data = search.LinearSearch(array1, number, &counter)
	elapsed = time.Since(start)

	fmt.Printf("Linear Search: Execution time: %s - Counter: %d\n", elapsed, counter)
	fmt.Printf("Number found: %d", data)
	fmt.Printf("\n\n")
}