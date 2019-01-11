package main

import "testing"

func TestExerciseReader(t *testing.T) {
	want := []byte{65, 65, 65, 65, 65}
	m := MyReader{}
	s := []byte{72, 55, 66, 88, 15}
	m.Read(s)

	for i, v := range want {
		if v != s[i] {
			t.Error("not same")
		}
	}
}
