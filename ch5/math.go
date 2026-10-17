package ch5

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	// More variables.
	i1, i2, i3 := 1, 2, 3
	// Addition.
	total := i1 + i2 + i3
	fmt.Println("total:", total)
	// Use math/big for higher-precision arithmetic
	// all the variables are big.Float type, so we can use the Add method to add them together. The Add method takes two big.Float pointers as arguments and returns a pointer to the result. We can chain the Add method calls to add multiple numbers together.
	var b1, b2, b3, bigSum big.Float
	// float example
	b1.SetFloat64(23.4)
	b2.SetFloat64(13.4)
	b3.SetFloat64(53.4)
	// Add the three big.Float numbers together using the Add method and store the result in bigSum. The Add method takes two big.Float pointers as arguments and returns a pointer to the result. We can chain the Add method calls to add multiple numbers together.
	bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
	fmt.Printf("Big sum = %.10g\n", &bigSum)

	// Example of Circle radius calculation
	CircleRadius := 15.6
	// Calculate the circumference of a circle using the formula: Circumference = 2 * π * r
	circumference := 2 * CircleRadius * math.Pi
	fmt.Printf("Circle circumference = %.2f\n", circumference)
}
