package main

import (
	"io"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)
	for i := 0; i < len(p); i++ {
		p[i] = p[i] + 1
		// fmt.Printf("p[%v] = %q\n", i, p[i])
	}

	return len(p), err
}

func ConvertAcstoInt(s *strings.Reader) []byte {
	c := make([]byte, s.Len())
	for {
		_, err := s.Read(c)
		if err == io.EOF {
			break
		}
	}
	return c
}

// func main() {
// 	s := strings.NewReader("ABCDEFGH")
// 	b := ConvertAcstoInt(s)
// 	// fmt.Println(b)
// 	// fmt.Println(b)
// 	// fmt.Println("-----------------")
// 	rot := rot13Reader{s}
// 	rot.Read(b)
// 	fmt.Println(string(b))

// }
