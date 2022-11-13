package main

import (
	"fmt"
	"net"
	"os"
)

// go run main.go www.google.com
// get an ip address from domain
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		fmt.Println("Usage: ", os.Args[0], "hostname")
		os.Exit(1)
	}
	name := os.Args[1]
	addr, err := net.ResolveIPAddr("ip", name)
	// addr, err := net.ResolveIPAddr("ip6", name) // ipv6 addr
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}

	fmt.Println("Resolved address is ", addr.String())
	os.Exit(0)
}
