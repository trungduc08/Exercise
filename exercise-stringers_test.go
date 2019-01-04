package main

import (
	"testing"
)

// type ListInput map[int]ListIP
// type ListIP map[string]IPAddr

func TestStringers(t *testing.T) {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	src := ConvertIPAdress(hosts)
	want := "loopback: 127.0.0.1\ngoogleDNS: 8.8.8.1"
	if src != want {
		t.Errorf(src, "\n", want)
	}
	// lstInput := ListInput{
	// 	0: ListIP{
	// 		"host1": [4]byte{127, 0, 0, 1},
	// 		"host2": [4]byte{127, 0, 0, 4},
	// 	},
	// 	1: ListIP{
	// 		"host3": [4]byte{127, 0, 0, 2},
	// 		"host4": [4]byte{127, 0, 0, 3},
	// 	},
	// }

	// lstOuput := ListInput{
	// 	0: ListIP{
	// 		"host1": [4]byte{127, 0, 0, 1},
	// 		"host2": [4]byte{127, 0, 0, 4},
	// 	},
	// 	1: ListIP{
	// 		"host3": [4]byte{127, 0, 0, 2},
	// 		"host4": [4]byte{127, 0, 0, 3},
	// 	},
	// }
	// for i, v := range lstInput {

	// }

	// input := map[string]IPAddr{
	// 	"host1": [4]byte{127, 0, 0, 1},
	// 	"host2": [4]byte{127, 0, 0, 2},
	// }
	// scr := ConvertIPAdress(input)
	// want := "host1: 127.0.0.1\nhosts2: 127.0.0.2"
	// if scr != want {
	// 	fmt.Errorf("not same")
	// }

}
