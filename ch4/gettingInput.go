package main

import (
	"fmt"
)
func main(){
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
fmt.Printf("OUTPUT:%v\n",s)
fmt.Println("--------=====-------") 
// end of program---------------------

}