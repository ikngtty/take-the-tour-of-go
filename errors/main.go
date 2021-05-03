package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func Sqrt(x float64) (float64, error) {
	// Validate
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	// Compute
	z := 1.0
	delta := 1.0
	for i := 1; math.Abs(delta) >= 1e-15 && i <= 1e5; i++ {
		fmt.Println(z)
		delta = (z*z - x) / (2 * z)
		z -= delta
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
