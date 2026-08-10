package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){

fmt.Println();
// import "bufio" and "os" to read input from console
reader:=bufio.NewReader(os.Stdin)
// display message to user
fmt.Print("read:")

// read string from console
str , _ := reader.ReadString('\n')
// out - > str
fmt.Println(str)

}