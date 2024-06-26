package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original Slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create Copy with only needed numbers
	needed_numbers := numbers[:len(numbers)-10]
	numbers_copy := make([]int, len(needed_numbers))
	copy(numbers_copy, needed_numbers)

	fmt.Printf("numbersCopy = %v\n", numbers_copy)
	fmt.Printf("length = %d\n", len(numbers_copy))
	fmt.Printf("capacity = %d\n", cap(numbers_copy))
}
