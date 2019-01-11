package main

import (
	"image/color"
	"testing"
)

func TestExerciseImage(t *testing.T) {

	img := Image{}
	w, h := 3, 2
	want := make([][]color.Color, w)
	input := ChangeColorInPixel(img, w, h)
	for i, _ := range want {
		want[i] = make([]color.Color, h)
	}
	want[0][0] = color.RGBA{
		R: 0,
		G: 0,
		B: 255,
		A: 255,
	}
	want[0][1] = color.RGBA{
		R: 1,
		G: 1,
		B: 255,
		A: 255,
	}
	want[1][0] = color.RGBA{
		R: 1,
		G: 1,
		B: 255,
		A: 255,
	}
	want[1][1] = color.RGBA{
		R: 2,
		G: 2,
		B: 255,
		A: 255,
	}
	want[2][0] = color.RGBA{
		R: 2,
		G: 2,
		B: 255,
		A: 255,
	}
	want[2][1] = color.RGBA{
		R: 3,
		G: 3,
		B: 255,
		A: 255,
	}

	if !Compare(input, want) {
		t.Error("not same")
	}

}
func Compare(a, b [][]color.Color) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		for j, k := range a[i] {
			if k != b[i][j] {
				return false
			}
		}
	}
	return true
}
