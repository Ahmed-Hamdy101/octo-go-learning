package ch5

import (
	"fmt"
	"octo-go-learning/ch5/modules"
)

func main() {
	// Use the BigSum function from the modules package.
	BigSum := modules.BigSum
	// Use the Circumference function from the modules package.
	Circumference := modules.Circumference
	// More variables.
	i1, i2, i3 := 1, 2, 3
	fmt.Println("Sum of 1, 2, 3 is:", BigSum(i1, i2, i3))
	fmt.Println("Circumference of radius 5 is:", Circumference(5))
}
