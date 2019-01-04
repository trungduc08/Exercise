package main

import "testing"

func TestFibonaci(t *testing.T) {
	x := []int{0, 1, 1, 2, 3}
	y := PrintFiBonaci(5)
	if !CompareSlice(x, y) {
		t.Error(x, y)
	}
}
func CompareSlice(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}
