package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	address := &numbers

	fmt.Printf("Original address: %p\n", address)
	fmt.Printf("Original slice: %v\n", numbers)

	// Adding a new element
	numbers = append(numbers, 6)
	numbers = append(numbers, 6)
	numbers = append(numbers, 6)
	numbers = append(numbers, 6)

	numbers = append(numbers, 6)
	numbers = append(numbers, 6)

	for i := 1; i <= 1000; i++ {
		numbers = append(numbers, 6)
	}

	fmt.Printf("Modified address: %p\n", &numbers)
	fmt.Printf("Modified slice: %v\n", numbers)
}
