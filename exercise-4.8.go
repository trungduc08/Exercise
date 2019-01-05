package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "trungduc123 tr"
	output := CountLetterByUnicodeLetter(s)
	fmt.Println(output)
}
func CountLetterByUnicodeLetter(s string) map[string]int {
	result := map[string]int{}
	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			result[fmt.Sprintf("%c", v)]++
		}
	}
	return result
}
