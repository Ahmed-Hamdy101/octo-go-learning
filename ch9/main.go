package main

import (
	"fmt"

	"octo-go-learning/ch9/arraypkg"
)

func main() {
	// inital new array with values
	keyboard := []string{"a", "b", "c", "d", "e", "f"}

	// store function in variables
	result := arraypkg.ArraySlice(keyboard, 2, len(keyboard)-1)
	// print full array
	fmt.Println("Original:", keyboard)
	// print subset
	fmt.Println("Subset:", result)
}
