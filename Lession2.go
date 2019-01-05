package main

import (
	"fmt"
	"sync"
	"time"
)

type I interface {
	M()
}
type T struct {
	S string
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}
func Describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func Do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v x 2 = %v\n", v, v*2)
	case string:
		fmt.Printf("%T\n", i)
	default:
		fmt.Printf("default")
	}
}

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v %v", p.name, p.age)
}

type MyError struct {
	when time.Time
	what string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.when, e.what)
}
func run() error {
	return &MyError{
		time.Now(),
		"did not work",
	}
}
func SumGo(s []int, c chan int) chan int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
	// fmt.Println(c)
	return c
}
func GoFibonaci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}

}

type SafeController struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeController) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}
func (c *SafeController) Value(key string) int {
	c.mux.Lock()
	return 1

}

var x int = 0

func Add(m *sync.Mutex) {
	m.Lock()
	x += 1
	m.Unlock()
}
func numbers(m *sync.Mutex) {
	m.Lock()
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}
	m.Unlock()

}
func alphabet(m *sync.Mutex) {
	m.Lock()
	for i := 'a'; i < 'f'; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%c\n", i)
	}
	m.Unlock()
}

// func main() {
// 	// var m sync.Mutex
// 	// go numbers(&m)
// 	// go alphabet(&m)
// 	// time.Sleep(10000 * time.Millisecond)
// 	// fmt.Println("finish")

// 	naturals := make(chan int)
// 	squares := make(chan int)
// 	// Counter
// 	go func() {
// 		for x := 0; x < 8; x++ {
// 			naturals <- x
// 		}
// 		close(naturals)
// 	}()
// 	// Squarer
// 	go func() {
// 		for {
// 			x, ok := <-naturals
// 			if !ok {
// 				break
// 			}
// 			squares <- x * x
// 		}
// 		close(squares)
// 	}()
// 	// Printer (in main goroutine)
// 	for value := range squares {
// 		fmt.Println(value)
// 	}
// }
type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

// func main() {
// 	// seen := make(map[string]bool)
// 	// input := bufio.NewScanner(os.Stdin)
// 	// for input.Scan() {
// 	// 	line := input.Text()
// 	// 	if !seen[line] {
// 	// 		seen[line] = true
// 	// 		fmt.Println(line)
// 	// 	}
// 	// }
// 	// if err := input.Err(); err != nil {
// 	// 	fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
// 	// 	os.Exit(1)
// 	// }
// 	// m = map[string]int{
// 	// 	"duc":  0,
// 	// 	"phuc": 1,
// 	// }
// 	// s := []string{}
// 	// for i, _ := range m {
// 	// 	s = append(s, i)
// 	// }

// 	// fmt.Println(k(s))
// 	// Add1(s)
// 	// fmt.Println(Count(s))

// 	counts := make(map[rune]int)
// 	// counts of Unicode characters
// 	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
// 	invalid := 0
// 	// count of invalid UTF-8 characters
// 	in := bufio.NewReader(os.Stdin)

// 	for {
// 		r, n, err := in.ReadRune() // returns rune, nbytes, error
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
// 			os.Exit(1)
// 		}
// 		if r == unicode.ReplacementChar && n == 1 {
// 			invalid++
// 			continue
// 		}
// 		counts[r]++
// 		utflen[n]++
// 	}
// 	fmt.Printf("rune\tcount\n")
// 	for c, n := range counts {
// 		fmt.Printf("%q\t%d\n", c, n)
// 	}
// 	fmt.Print("\nlen\tcount\n")
// 	for i, n := range utflen {
// 		if i > 0 {
// 			fmt.Printf("%d\t%d\n", i, n)
// 		}
// 	}
// 	if invalid > 0 {
// 		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
// 	}

// }

var m = make(map[string]int)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}
func Add1(list []string) {
	m[k(list)]++
}
func Count(list []string) int {
	return m[k(list)]
}

type Employee struct {
	id      int
	name    string
	address string

	position  string
	salary    int
	managerId int
}

var v []Employee

// func main() {
// 	v = []Employee{
// 		Employee{
// 			id:        0,
// 			name:      "duc",
// 			address:   "hn",
// 			position:  "center",
// 			salary:    12,
// 			managerId: 13,
// 		},
// 		Employee{
// 			id:        1,
// 			name:      "phuc",
// 			address:   "hn",
// 			position:  "left",
// 			salary:    20,
// 			managerId: 25,
// 		},
// 	}
// 	fmt.Println(EmployeeByID(12))

// }

func EmployeeByID(id int) *Employee {
	temp := -1

	for _, val := range v {
		if val.id == id {
			temp = id
			break
		}

	}
	if temp < 0 {
		return nil
	}
	return &(v[temp])
}
