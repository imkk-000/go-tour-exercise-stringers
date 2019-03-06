package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ipAddr IPAddr) String() string {
	var ipString string
	for index, ipElement := range ipAddr {
		ipString += strconv.Itoa(int(ipElement))
		if index < 3 {
			ipString += "."
		}
	}
	return ipString
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
