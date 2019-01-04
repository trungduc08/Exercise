package main

import (
	"math"
	"testing"
)

type ErrorTestFail struct {
	defaultValue float64
	err          string
}
type ErrorTestPass struct {
	value float64
	err   error
}

func TestExerciseErrors(t *testing.T) {
	// failed
	// var inputNumberFail float64 = -8
	// val, err := Sqrt(inputNumberFail)
	// errFail := ErrorTestFail{
	// 	0,
	// 	fmt.Sprintf("cannot Sqrt negative number:%v", int(inputNumberFail)),
	// }
	// want := fmt.Sprintf("%v ", int(errFail.defaultValue)) + errFail.err
	// src := fmt.Sprintf("%v ", int(val)) + fmt.Sprint(err)
	// if want != src {
	// 	t.Error("not same")
	// }

	//pass
	var inputNumberRight float64 = 9
	errorPass := ErrorTestPass{
		math.Sqrt(inputNumberRight),
		nil,
	}
	errorPassWant := ErrorTestPass{}
	valRight, errRight := Sqrt(inputNumberRight)
	errorPassWant.value = valRight
	errorPassWant.err = errRight
	if errorPass != errorPassWant {
		t.Error("not same")
	}

}
