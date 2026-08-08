package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// import "bufio" and "os" to read input from console
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("real point(x):")
	// create string + new getInput
	str, _ := reader.ReadString('\n')
	// convert string to float
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	//  if err is not nothing then print float , go to if exec print error
	if err != nil {
		fmt.Println("float", f)
	} else {
		fmt.Println(err)
	}

	// out - > str
	fmt.Println(str)

}
