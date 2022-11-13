package main

import (
	"fmt"
	"net"
	"os"
)

// example:
// go run main.go 127,0,0,1 => 127,0,0,1 // ipv4
// go run main.go 0:0:0:0:0:0:0:1 => ::1 // ipv6
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}
