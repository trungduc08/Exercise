package main

import (
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
func CountBit(x [32]byte) int {
	sum := 0
	for _, v := range x {
		sum += PopCount(uint64(v))
	}
	return sum
}

// func main() {
// 	count := 0
// 	c1 := sha256.Sum256([]byte("A"))
// 	c2 := sha256.Sum256([]byte("X"))
// 	src1 := []string{}
// 	src2 := []string{}

// 	for i, v := range c1 {
// 		src1 = append(src1, fmt.Sprintf("%v", strconv.FormatUint(uint64(v), 2)))
// 		src2 = append(src2, fmt.Sprintf("%v", strconv.FormatUint(uint64(c2[i]), 2)))
// 	}
// 	// fmt.Println(src1)
// 	// fmt.Println(src2)
// 	for i, v := range src1 {
// 		count = count + CountDifferent(fmt.Sprintf("%s", v), fmt.Sprintf("%s", src2[i]))
// 	}
// 	fmt.Println(count)

// }
func CountDifferent(a, b string) int {
	count := 0
	if len(a) == len(b) {
		for i, v := range a {
			if fmt.Sprintf("%c", v) != fmt.Sprintf("%c", b[i]) {
				count = count + 1
			}
		}
	} else if len(a) < len(b) {
		for i, v := range a {
			if fmt.Sprintf("%c", v) != fmt.Sprintf("%c", b[i]) {
				count++
			}
		}
		count = count + len(b) - len(a)
	} else {
		for i, v := range b {
			if fmt.Sprintf("%c", v) != fmt.Sprintf("%c", a[i]) {
				count++
			}
		}
		count = count + len(a) - len(b)
	}
	return count
}
