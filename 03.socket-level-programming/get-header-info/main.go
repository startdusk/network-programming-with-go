package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

// go run . www.baidu.com:80
// Output:
// HTTP/1.0 302 Found
// Content-Length: 17931
// Content-Type: text/html
// Date: Tue, 25 Oct 2022 15:12:31 GMT
// Etag: "54d9748e-460b"
// Server: bfe/1.0.8.18

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := io.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
