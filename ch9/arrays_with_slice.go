package main

import "fmt"

func main() {
	// create a slice of strings
	var keyboard []string
	// add values to the slice
	keyboard = append(keyboard, "a", "b", "c")
	fmt.Println(keyboard)

	// add more values to the slice
	keyboard = append(keyboard, "d", "e", "f")
	fmt.Println(keyboard)

	// create a subset of the slice  #def
	keyboard = keyboard[2 : len(keyboard)-1]

	// slice from index 2 to the end of the slice
	fmt.Println("Subset:", keyboard)

	// Make a slice with length and capacity of 50.
	button := make([]int, 50)
	// fill the slice with values
	for index := 0; index < 10; index++ {
		button[index] = index + 1
	}

	// print the slice and its capacity
	button = append(button, 1000, 9000)
	fmt.Println(button)

	// The exact capacity after append is chosen by Go's runtime.
	fmt.Println("Capacity:", cap(button))
}
