package main

import (
	"fmt"
	"math"
)

const EPSILON = 0.0000000001

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {

	return fmt.Sprintf("cannot Sqrt negative number:%v", float64(e))
}
func Sqrt(x float64) (float64, error) {
	if x < 0 {

		return 0, ErrNegativeSqrt(x)
	}
	return CalSqrt(x), nil
}
func CalSqrt(x float64) float64 {
	z := float64(1)
	for math.Abs(z*z-x) >= EPSILON {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

// type ErrorTestFail struct {
// 	defaultValue float64
// 	err          string
// }
// type ErrorTestPass struct {
// 	value float64
// 	err   error
// }

// func main() {
// 	var inputNumberRight float64 = 9
// 	errorPass := ErrorTestPass{
// 		math.Sqrt(inputNumberRight),
// 		nil,
// 	}
// 	errorPassWant := ErrorTestPass{}
// 	valRight, errRight := Sqrt(inputNumberRight)
// 	errorPassWant.value = valRight
// 	errorPassWant.err = errRight

// 	fmt.Println(errorPass)
// 	fmt.Println()
// 	fmt.Println(errorPassWant)
// }
