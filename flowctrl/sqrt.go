package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		diff := (z * z - x) / (2 * z)
		fmt.Println(i, z, diff)
		if math.Abs(diff) < 0.000000000000001 {
			return z
		}
		z -= diff
	}
	return z
}

func main() {
	const num = 5
	fmt.Println(Sqrt(num))
	fmt.Println("math.Sqrt(num): ", math.Sqrt(num))
}