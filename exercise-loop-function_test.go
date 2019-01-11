package main

import (
	"math"
	"testing"
)

func TestLoopAndFunction(t *testing.T) {
	want := math.Sqrt(8)
	input := SqrtAddCoditional(8)
	if math.Abs(float64(want)-input) >= EPSILON {
		t.Error(want, "\n", input)
	}

}
