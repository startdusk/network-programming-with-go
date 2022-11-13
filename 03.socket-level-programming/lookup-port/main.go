package main

import (
	"fmt"
	"net"
	"os"
)

// go run . tcp telnet
// Output:
// Service port  23
// go run . tcp http
// Output:
// Service port  80
func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr,
			"Usage: %s network-type service\n",
			os.Args[0])
		os.Exit(1)
	}
	networkType := os.Args[1]
	service := os.Args[2]
	port, err := net.LookupPort(networkType, service)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}
	fmt.Println("Service port ", port)
	os.Exit(0)
}
