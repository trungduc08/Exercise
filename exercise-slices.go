package main

func Pic(dx, dy int) [][]uint8 {
	// dx : number of column
	// dy: number of row
	a := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		a[i] = make([]uint8, dx)

	}

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {

			a[y][x] = (uint8)(x + y)
		}
	}

	return a

}

// func main() {

// 	var src string
// 	src = pic.Show(Pic)
// 	fmt.Println(src)

// }
