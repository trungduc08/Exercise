package main

import (
	"image"
	"image/color"
)

type Image struct {
	width, height int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{(uint8)(x + y), (uint8)(x + y), 255, 255}
}

func ChangeColorInPixel(img Image, w int, h int) [][]color.Color {
	img = Image{w, h}
	img.Bounds()
	s := make([][]color.Color, w)

	for i := 0; i < w; i++ {
		s[i] = make([]color.Color, h)

	}
	for i := img.Bounds().Min.X; i < img.Bounds().Max.X; i++ {
		for j := img.Bounds().Min.Y; j < img.Bounds().Max.Y; j++ {
			img.At(i, j)
			// fmt.Printf("%+v", img.At(i, j))
			s[i][j] = img.At(i, j)
			// fmt.Printf("%+v", s[i][j])

			// s[i][j] = color.RGBA{
			// 	R: 0,
			// 	G: 255,
			// 	B: 255,
			// 	A: 255,
			// }

		}
	}
	return s
}

// func main() {

// 	w, h := 3, 2
// 	img := Image{}
// 	want := make([][]color.Color, w)
// 	input := ChangeColorInPixel(img, w, h)
// 	for i, _ := range want {
// 		want[i] = make([]color.Color, h)
// 	}
// 	want[0][0] = color.RGBA{
// 		R: 0,
// 		G: 0,
// 		B: 255,
// 		A: 255,
// 	}
// 	want[0][1] = color.RGBA{
// 		R: 1,
// 		G: 1,
// 		B: 255,
// 		A: 255,
// 	}
// 	want[1][0] = color.RGBA{
// 		R: 1,
// 		G: 1,
// 		B: 255,
// 		A: 255,
// 	}
// 	want[1][1] = color.RGBA{
// 		R: 2,
// 		G: 2,
// 		B: 255,
// 		A: 255,
// 	}
// 	want[2][0] = color.RGBA{
// 		R: 2,
// 		G: 2,
// 		B: 255,
// 		A: 255,
// 	}
// 	want[2][1] = color.RGBA{
// 		R: 3,
// 		G: 3,
// 		B: 255,
// 		A: 255,
// 	}
// 	// for i, _ := range want {
// 	// 	for j, k := range want[i] {
// 	// 		if k != input[j]{
// 	// 			fmt.Println("false")
// 	// 		}
// 	// 	}
// 	// }
// 	// fmt.Println(Compare(want, input))

// 	// fmt.Println()
// 	// fmt.Println(ChangeColorInPixel(img, w, h))

// 	// img := Image{}
// 	// w, h := 3, 2
// 	// s := ChangeColorInPixel(img, w, h)
// 	// fmt.Println(s)
// 	// s := []int{1, 2, 3, 4}
// 	// fmt.Println(img.At(0, 0))
// 	// fmt.Println(s)

// 	// for i, _ := range s {
// 	// 	for j, _ := range s[i] {
// 	// 		fmt.Print(s[i][j])
// 	// 	}
// 	// }
// 	// img := Image{3, 2}
// 	// img.Bounds()
// 	// s := make([][]color.Color, 3)
// 	// for i := 0; i < 3; i++ {
// 	// 	s[i] = make([]color.Color, 2)

// 	// }
// 	// for i := img.Bounds().Min.X; i < img.Bounds().Max.X; i++ {
// 	// 	for j := img.Bounds().Min.Y; j < img.Bounds().Max.Y; j++ {
// 	// 		img.At(i, j)
// 	// 		// fmt.Println(img.At(i, j))
// 	// 		s[i][j] = img.At(i, j)

// 	// 	}
// 	// }
// 	// fmt.Println(s)

// 	// for _, v := range s {
// 	// 	fmt.Println(v)
// 	// }
// }

// func Compare(a, b [][]color.Color) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}
// 	for i, _ := range a {
// 		for j, k := range a[i] {
// 			if k != b[i][j] {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }
