package main

import "fmt"

func main() {
	s := "abcxyz"
	input := ConvertStringToArray(s)
	ReverseString(input)
	output := ConvertArrayToString(input)
	fmt.Println(output)

}
func ConvertStringToArray(s string) []string {
	result := []string{}
	for _, v := range s {
		result = append(result, fmt.Sprintf("%c", v))
	}
	return result
}
func ConvertArrayToString(src []string) string {
	result := ""
	for _, v := range src {
		result = result + v
	}
	return result
}
func ReverseString(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
