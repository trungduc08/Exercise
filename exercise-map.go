package main

import (
	"fmt"
	"strings"
)

var mm1 = map[string]int{
	"Bell Labs": 40,
	"Google":    37,
}

func WordCount(s string) map[string]int {

	var mm = map[string]int{}
	a := strings.Fields(s)
	// fmt.Println(a)
	// fmt.Println(mm)
	// fmt.Println(len(a))
	for _, v := range a {
		// fmt.Println(mm[v])
		mm[v]++
		// fmt.Println(mm[v])
	}
	// fmt.Println(a)
	fmt.Println(mm)
	return mm
}

// func main() {

// 	// fmt.Println(mm1["Google"])
// 	WordCount("trung duc trung du")
// 	// var a = []int{1, 2, 4, 8, 16, 32, 64}
// 	// for i, v := range a {
// 	// 	fmt.Println("2^", i, "=", v)
// 	// }

// }
