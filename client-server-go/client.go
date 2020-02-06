package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// conn, err := net.Dial("tcp", "golang.org:80")
	ipandport := "localhost:8080" // or "127.0.0.1:8080"
	conn, err := net.Dial("tcp", ipandport)

	if err != nil {
		// handle error

	}

	reader := bufio.NewReader(os.Stdin) // read in input from stdin
	fmt.Print(">Client Side| Please type a message to send: ")
	message, _ := reader.ReadString('\n')

	// Tranmissting message to server
	fmt.Fprintf(conn, message+"\n")

	// waiting to recv reply from the server
	messagercv, _ := bufio.NewReader(conn).ReadString('\n')

	fmt.Println(">Client Side| Message from the server :" + messagercv)

}
