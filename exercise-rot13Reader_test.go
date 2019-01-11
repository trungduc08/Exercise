package main

import (
	"strings"
	"testing"
)

func TestExerciseRot13Reader(t *testing.T) {
	want := "BCDEFGHI"

	s := strings.NewReader("ABCDEFGH")
	b := ConvertAcstoInt(s)
	rot := rot13Reader{s}
	rot.Read(b)

	input := string(b)

	if want != input {
		t.Error("not same")
	}
}
