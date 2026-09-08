package main

import (
	"fmt"
)

func main() {
	// pointer in GO for calling pointer use strict  as * name var
	// define * name var
	var p *int

	if p != nil {
		fmt.Println("p is not nil", *p)
	} else {
		fmt.Println("p is nil")
	}
	var v int = 42
	p = &v
	if p != nil {
		fmt.Println("p is not nil", *p)
	} else {
		fmt.Println("p is nil")
	}
	//& for bind the value in pointer as memmory
	var val float64 = 45.3
	pointer1 := &val
	*pointer1 = *pointer1 / 32
	fmt.Println("pointer1 as new pointer:", *pointer1)
	fmt.Println("val as new pointer:", val)

}
