package main

import "fmt"

func main() {
	var animals [5]string
	animals[0] = "lion"
	animals[1] = "tiger"
	animals[2] = "wolf"
	animals[3] = "cow"
	animals[4] = "elephant"
	fmt.Println(animals)

	// intial arrays with value
	var goos = [2]int{9999, 2112}
	fmt.Println(goos)
	fmt.Println(len(animals))

	fmt.Println(Swap("friend"))
}

// example of array  swap I to E with word friend
func Swap(s string) string {
	// convert string to array
	words := []rune(s)
	words[2], words[3] = words[3], words[2]
	return string(words)
}
