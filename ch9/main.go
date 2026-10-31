package main

import (
	"fmt"

	ch9 "ch9/arraypkg"
)

func main() {

	keyboard := []string{"a", "b", "c", "d", "e", "f"}

	result := ch9.ArraySlice(keyboard, 2, len(keyboard)-1)
	fmt.Println("Original:", keyboard)
	fmt.Println("Subset:", result)
}
