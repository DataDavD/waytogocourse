package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	var (
		host          = "www.apache.org"
		port          = "80"
		remote        = host + ":" + port
		msg    string = "GET / \n"
		data          = make([]uint8, 4096)
		read          = true
		count         = 0
	)
	// create the socket
	conn, err := net.Dial("tcp", remote)
	// send our message. an HTTP GET request in this case
	io.WriteString(conn, msg)
	// read the response from the webserver
	for read {
		count, err = conn.Read(data)
		read = err == nil
		fmt.Printf(string(data[0:count]))
	}
	conn.Close()
}
