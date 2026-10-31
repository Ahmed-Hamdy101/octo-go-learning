package main

import (
	"fmt"

	"octo-go-learning/ch9/arraypkg"
)

func main() {
	// Create a slice of strings.
	keyboard := []string{"a", "b", "c", "d", "e", "f"}

	// Call the reusable slice function and store the result.
	result := arraypkg.ArraySlice(keyboard, 2, len(keyboard)-1)

	// Print the original list and the sliced result.
	fmt.Println("Original:", keyboard)
	fmt.Println("Subset:", result)
}
