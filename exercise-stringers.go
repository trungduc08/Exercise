package main

import (
	"fmt"
)

type IPAddr [4]byte

func ConvertIPAdress(_map map[string]IPAddr) string {
	var src string
	for i, _ := range _map {
		if src == "" {
			src = fmt.Sprintf("%v%v: %v.%v.%v.%v", src, i, _map[i][0], _map[i][1], _map[i][2], _map[i][3])
		} else {
			src = fmt.Sprintf("%v\n%v: %v.%v.%v.%v", src, i, _map[i][0], _map[i][1], _map[i][2], _map[i][3])
		}

	}
	return src
}

type ListInput map[int]ListIP
type ListIP map[string]IPAddr

// func main() {

// 	hosts := map[string]IPAddr{
// 		"loopback":  {127, 0, 0, 1},
// 		"googleDNS": {8, 8, 8, 8},
// 	}
// 	want := "loopback: 127.0.0.1\ngoogleDNS: 8.8.8.8"
// 	input := ConvertIPAdress(hosts)
// 	// fmt.Println(hosts)
// 	fmt.Println(input)
// 	fmt.Println(want)
// 	if !CompareString(want, input) {
// 		fmt.Println("fa")
// 	} else {
// 		fmt.Println("true")
// 	}
// 	// for i, _ := range hosts {
// 	// 	fmt.Printf("%v: %v\n", i, hosts[i])
// 	// }

// }
func CompareString(s1, s2 string) bool {
	if s1 != s2 {
		return false
	}
	return true

}
