package main

import "fmt"

func main() {
	// create a slice of strings
	var keyboard []string
	// add values to the slice
	keyboard = append(keyboard, "a", "b", "c")
	fmt.Println(keyboard)
	// add more values to the slice
	keyboard = append(keyboard , "d","e","f")
	fmt.Println(keyboard)
	// create a subset of the slice
	keyboard = keyboard[2 : len(keyboard)-1] // #1
	// keyboard = append(keyboard[1:len(keyboard)-1]) // #2
	fmt.Println("Subset:",keyboard)

// ------------------Make data types-------------------
	// make a slice of integers with length 10 and capacity 10
	button := make([]int,50,50)
	 button[0] = 1
	 button[1] = 2
	 button[2] = 3
	 button[3] = 4
	 button[4] = 5
	 button[5] = 6
	 button[6] = 7
	 button[7] = 8
	 button[8] = 9
	 button[9] = 10
	button = append(button ,1000,9000) 
	fmt.Println(button)
	// capacity of the slice 
	/**
	*  Correct. ⚠ You cannot rely on the exact capacity after append.
	You can rely on:
	But the new capacity might be 100, 112, 128, or another value depending on Go’s runtime and architecture.
	If you need a predictable capacity, allocate it yourself
	*/
	fmt.Println("Capacity:", cap(button))
}
