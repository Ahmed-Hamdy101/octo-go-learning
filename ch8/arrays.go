package main

import "fmt"

func main() {
	var animals [5]string
	animals[0] = "lion"
	animals[1] = "tiger"
	animals[2] = "wolf"
	animals[3] = "cow"
	animals[4] = "elphant"
	fmt.Println(animals)
	
	// intial arrays with value
	var goos = [2]int{9999,2112}
	fmt.Println(goos)
}