package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

// func main() {
// 	// ch1 := sha512.Sum512([]byte("Hello"))

// 	// fmt.Printf("%x\n", ch1)
// 	// fmt.Println(ConvertSha512("Hello"))

// 	var src, input string
// 	fmt.Println("Nhap chuoi")
// 	fmt.Scan(&src)
// 	fmt.Println("1. sha512 \t 2. sha382 \t 3. sha256")
// 	fmt.Println("Nhap lua chon")
// 	fmt.Scan(&input)
// 	switch input {
// 	case "1":
// 		fmt.Println(ConvertSha512(input))
// 	case "2":
// 		fmt.Println(ConvertSha384(input))
// 	case "3":
// 		fmt.Println(ConvertSha256(input))
// 	default:
// 		fmt.Println("Finish")

// 	}

// }

func ConvertSha512(s string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(s)))
}
func ConvertSha384(s string) string {
	return fmt.Sprintf("%x", sha512.Sum384([]byte(s)))
}
func ConvertSha256(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}
