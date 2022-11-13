package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

// For an IPv4 address of 103.232.159.187 on a /24 network, we get the following:
// go run main.go 103.232.159.187 24 32
// For an IPv6 address fda3:97c:1eb:fff0:5444:903a:33f0:3a6b where the network component is
// fda3:97c:1eb, the subnet is fff0, and the device part is 5444:903a:33f0:3a6b, we get the following:
// go run main.go fda3:97c:1eb:fff0:5444:903a:33f0:3a6b 52 128
func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr ones bits\n", os.Args[0])
		os.Exit(1)
	}

	dotAddr := os.Args[1]
	ones, _ := strconv.Atoi(os.Args[2])
	bits, _ := strconv.Atoi(os.Args[3])

	addr := net.ParseIP(dotAddr)
	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}

	mask := net.CIDRMask(ones, bits)
	network := addr.Mask(mask)
	fmt.Println("Address is ", addr.String(),
		"\nMask length is ", bits,
		"\nLeading ones count is ", ones,
		"\nMask is (hex) ", mask.String(),
		"\nNetwork is ", network.String())
	os.Exit(0)
}
