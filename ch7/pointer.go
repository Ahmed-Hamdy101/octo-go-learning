package main

import (
	"fmt"
)

func main() {
	// pointer in GO for calling pointer use strict  as * name var
	// define * name var
	var p *int

	// check if pointer is nil or not
	if p != nil {
		fmt.Println("p is not nil", *p)
	} else {
		fmt.Println("p is nil")
	}

	// define pointer with value
	var v int = 42
	// bind pointer with value
	//& for bind the value in pointer as memmory
	p = &v

	// check if pointer is nil or not
	if p != nil {
		fmt.Println("p is not nil", *p)
	} else {
		fmt.Println("p is nil")
	}

	// define variable with value
	var val float64 = 45.3
	// define pointer with value
	pointer1 := &val

	// pointer1 is a pointer to val, so *pointer1 is the value of val and vakl is the value of val
	fmt.Println("pointer1:", *pointer1)
	fmt.Println("val:", val)

	// change the value of val using pointer1
	*pointer1 = 45.414

	// print the value of val and pointer1
	fmt.Println("pointer1 as new pointer:", *pointer1)
	fmt.Println("val as new pointer:", val)

}
