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
