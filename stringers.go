package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	str := ""
	for _, b := range ip {
		str += fmt.Sprintf("%d", b) + "."
	}
	return str[:len(str)-1]
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
