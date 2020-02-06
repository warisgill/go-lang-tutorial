package clientserver

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	// get message, output
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message Received:", string(message))
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}
