package main

import (
	"fmt"
	"math/big"
)

func main() {
	// More variables.
	i1, i2, i3 := 1, 2, 3
	// Addition.
	total := i1 + i2 + i3
	fmt.Println("total:", total)
	// Use math/big for higher-precision arithmetic.
	var b1, b2, b3, bigSum big.Float
	// float
	b1.SetFloat64(23.4)
	b2.SetFloat64(13.4)
	b3.SetFloat64(53.4)
	bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
	fmt.Printf("Big sum = %s\n", bigSum.Text('g', 10))

}
