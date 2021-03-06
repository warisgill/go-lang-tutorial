package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

// do not import service struct

type Args struct {
	A, B int
}

type Answer struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server:port")
		os.Exit(1)
	}

	service := os.Args[1] // address (i.e., IP:port) of the service

	client, err := rpc.Dial("tcp", service) // connecting to the service

	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := Args{17, 10}
	var reply int
	err = client.Call("PS3.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Answer: %d*%d=%d\n", args.A, args.B, reply)

	var quot Answer
	err = client.Call("PS3.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}

	fmt.Printf("Answer: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

}
