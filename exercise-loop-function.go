package main

import (
	"math"
)

// const EPSILON = 0.0000000001

// func Sqrt(x float64) float64 {
// 	z := float64(x)
// 	for i := 0; i < 10; i++ {
// 		z -= (z*z - x) / (2 * z)
// 		fmt.Println(z)

// 	}
// 	return z
// }

func SqrtAddCoditional(x float64) float64 {
	z := float64(1)
	for math.Abs(z*z-x) >= EPSILON {
		z -= (z*z - x) / (2 * z)

	}
	return z
}

// func main() {
// 	fmt.Println(math.Sqrt(4))
// 	fmt.Println("--------------")
// 	fmt.Println(SqrtAddCoditional(4))

// }
