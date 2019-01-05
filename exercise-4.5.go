package main

import "fmt"

func main() {
	src := []string{
		"Hello",
		"Hello",
		"Word",
		"Lang",
		"Go",
		"Go",
		"Go",
		"Go",
		"Lasst",
	}
	result := map[int]Result{}
	result = RemoveEliminateAdjacentDuplicate(src)
	for i, _ := range result {
		for j := result[i].index; j < result[i].index+result[i].count; j++ {
			src[j] = ""
		}
	}
	lastResult := Nonempty(src)
	for _, v := range lastResult {
		fmt.Println(v)
	}
}

type Result struct {
	src   string
	index int
	count int
}

func RemoveEliminateAdjacentDuplicate(src []string) map[int]Result {
	result := map[int]Result{}
	j := 0
	count := 1
	for i := 0; i < len(src)-1; i++ {
		if src[i] == src[i+1] {
			count++
		} else {
			if count > 1 {
				result[j] = Result{src: src[i], index: i - count + 1, count: count}
				count = 1
				j++
			}
		}
		if i >= len(src)-2 {

			if count > 1 {
				result[j] = Result{src: src[i], index: i - count + 2, count: count}
				count = 1
				j++
			}
		}
	}
	return result
}
func Nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
