package main

import (
	"unicode"
)

type Stored struct {
	src   rune
	index int
	count int
}

// func main() {
// 	input := "  trung duc        trung        duc      "
// 	output := []rune{}
// 	runeSrc := []rune{}
// 	for _, v := range input {
// 		runeSrc = append(runeSrc, v)
// 	}
// 	// fmt.Println(runeSrc)
// 	lastRune := RemoveSpace(runeSrc)
// 	// fmt.Println(lastRune)
// 	for i, _ := range lastRune {
// 		if i == 0 {
// 			for j := lastRune[i].index; j < lastRune[i].index+lastRune[i].count; j++ {
// 				runeSrc[j] = 0
// 			}
// 		} else {
// 			for j := lastRune[i].index + 1; j < lastRune[i].index+lastRune[i].count; j++ {
// 				runeSrc[j] = 0
// 			}
// 		}
// 	}
// 	for _, v := range runeSrc {
// 		if v != 0 {
// 			output = append(output, v)
// 		}

// 	}
// 	fmt.Println(string(output))
// }
func RemoveSpace(src []rune) map[int]Stored {
	result := map[int]Stored{}
	j := 0
	count := 1
	for i := 0; i < len(src)-1; i++ {
		if unicode.IsSpace(src[i]) && unicode.IsSpace(src[i+1]) {
			count++
		} else {
			if count > 1 {
				result[j] = Stored{src: src[i], index: i - count + 1, count: count}
				count = 1
				j++
			}
		}
		if i >= len(src)-2 {
			if count > 1 {
				result[j] = Stored{src: src[i], index: i - count + 2, count: count}
				count = 1
				j++
			}
		}
	}
	return result
}
