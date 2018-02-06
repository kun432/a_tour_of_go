package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := 1.0
	for diff := 1.0; math.Abs(diff) > 1e-6; {
		diff = (z*z - x) / (2 * z)
		z -= diff
		fmt.Println("debug    : ", z)
	}
	return z
}

func main() {
	fmt.Println("sqrt     : ", sqrt(2))
	fmt.Println("math.Sqrt: ", math.Sqrt(2))
}
