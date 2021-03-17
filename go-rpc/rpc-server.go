package main

/*
	Package rpc provides access to the exported methods of an object across a network
	or other I/O connection.
	A server registers an object, making it visible as a service with the name of the type of the object.
	After registration, exported methods of the object will be accessible remotely.
*/

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}

type Answer struct {
	Quo, Rem int
}

type PS3 struct {
}

func (t *PS3) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *PS3) Divide(args *Args, reply *Answer) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	reply.Quo = args.A / args.B
	reply.Rem = args.A % args.B
	return nil
}

/*
	Package rpc provides access to the exported methods of an object across a
	network or other I/O connection.
	A server registers an object, making it visible as a service with the name
	of the type of the object. After registration, exported methods of the object
	will be accessible remotely.
	https://golang.org/pkg/net/rpc/
*/

func main() {

	// creating and regestring PS3 object as a Service on the network
	magic := new(PS3)
	rpc.Register(magic)

	// sarting server
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	fmt.Println(">Serivice is running.... ")

	// similar to server code
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
