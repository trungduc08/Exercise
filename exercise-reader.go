package main

import "fmt"

type MyReader struct {
}

func (r MyReader) Read(s []byte) (int, error) {
	for i, _ := range s {
		s[i] = 65
	}
	return len(s), nil
}

func main() {
	m := MyReader{}
	s := []byte{72, 55, 66, 88, 15}
	m.Read(s)
	for _, v := range s {
		fmt.Printf("%c ", v)
	}

}
