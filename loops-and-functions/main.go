package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	delta := 1.0
	for i := 1; math.Abs(delta) >= 1e-15 && i <= 1e5; i++ {
		fmt.Println(z)
		delta = (z*z - x) / (2 * z)
		z -= delta
	}
	return z
}

func main() {
	fmt.Println(Sqrt(42))
}
