package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	arr := RotateOnePass(s)
	fmt.Println(arr)

}
func RotateOnePass(input []int) []int {
	for i := 0; i < len(input)/2; i++ {
		temp := input[i]
		copy(input[i:], input[len(input)-i-1:len(input)-i])
		input[len(input)-1-i] = temp
	}
	return input[:len(input)]

}
