package main

import (
	"reflect"
	"testing"
)

func TestExerciseMap(t *testing.T) {
	want := map[string]int{
		"trung": 2,
		"duc":   1,
		"du":    1,
	}
	input := WordCount("trung duc trung du")
	if !reflect.DeepEqual(want, input) {
		t.Error("2 map not same")
	}
}
