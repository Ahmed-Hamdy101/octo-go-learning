package main

import (
	"fmt"
)

func yourAge() {
	// intial
	var s string
	// display message to user
	fmt.Println("---------------------")
	fmt.Println("how old are you:")
	fmt.Print("Enter:")
	// get input from user
	fmt.Scanln(&s)
	// display output -------------------
	fmt.Println("---------------------")
	fmt.Println("--------=OUTPUT=--------")
	fmt.Printf("Your age is: %v\n", s)
	fmt.Println("--------=====-------")
	// end of program---------------------

}
