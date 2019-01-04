package main

import (
	"testing"

	"golang.org/x/tour/tree"
)

func TestExerciseTree(t *testing.T) {

	want := true
	input := CheckSame(tree.New(1), tree.New(1))
	if want != input {
		t.Error("two tree not same")
	}

}
func CompareTwoChannel(a, b chan int) bool {

	if cap(a) != cap(b) {
		return false
	}
	for {
		i, valid1 := <-a
		j, _ := <-b
		if !valid1 {
			break
		}
		if i != j {
			return false

		}
	}
	return true
}
