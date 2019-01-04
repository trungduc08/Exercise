package main

import "testing"

func TestExerciseWeb(t *testing.T) {
	want := []string{
		"found: https://golang.org/ The Go Programming Language",
		"found: https://golang.org/pkg/ Packages",
		"not found: https://golang.org/cmd/",
		"found: https://golang.org/pkg/fmt/ Package fmt",
		"found: https://golang.org/pkg/os/ Package os",
	}

	input := ReadFile()
	if !CompareArrString(input, want) {
		t.Error("not same")
	}
}

func CompareArrString(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true

}
