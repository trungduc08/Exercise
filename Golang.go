package main

import (
	"math/cmplx"
)

func Swap(x, y int) (int, int) {
	return y, x
}
func Split(sum int) (x, y int) {
	x = sum + 2
	y = sum - 2
	return
}

var c, python, java = true, false, "no!"
var (
	toBe   bool       = false
	maxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	i      int        = 3
	str    string     = "hello"
)

const (
	big   = 1 << 20
	small = big >> 19
)

func NeedInt(x int) int {
	return x*10 + 1
}
func NeedFloat(x float64) float64 {
	return x * 0.1

}

// func main(){
// 	fmt.Println(Swap(4, 3))
// 	fmt.Println(NeedInt(small))

// 	fmt.Println(NeedFloat(big))

// 	v := 12 + 9i
// 	fmt.Printf(" v is of Type %T \n", v)

// 	fmt.Println("---------------------------")
// 	var x, y int = 3, 6
// 	var z float64 = math.Sqrt(float64(x*x + y*y))
// 	fmt.Printf("%v %v %v \n", x, y, z)
// 	fmt.Println("---------------------------");
// 	fmt.Printf("%T %v\n", toBe, toBe)
// 	fmt.Printf("%T %v\n", maxInt, maxInt)
// 	fmt.Printf("%T %v\n", z, z)
// 	fmt.Printf("%T %v\n", i, i)
// 	fmt.Printf("%T %v\n", str, str)
// 	fmt.Println("---------------------------");
//     rand.Seed(time.Now().UTC().UnixNano())
// 	fmt.Println(rand.Intn(10))
// 	fmt.Println(math.Sqrt(7))

// }
