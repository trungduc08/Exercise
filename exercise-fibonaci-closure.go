package main

func fibonacci() func() int {
	x, y := 1, 0
	var temp int
	return func() int {
		temp = x
		x = y
		y = y + temp
		return x
	}
}

func PrintFiBonaci(n int) (arr []int) {
	f := fibonacci()
	if n < 0 {
		return nil
	}
	for i := 0; i < n; i++ {
		arr = append(arr, f())
	}
	return arr

}

// func main() {
// 	fmt.Println(PrintFiBonaci(5))
// }
