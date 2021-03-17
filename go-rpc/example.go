package main

import (
	"errors"
	"fmt"
)

type Args struct {
	A, B int
}

type Answer struct {
	Quo, Rem int
}

type PS3 struct {
}

/*
	Go does not have classes. However, you can define methods on types.

	A method is a function with a special receiver argument.
*/

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

func main() {

	fmt.Printf("Simple Example, \n")
	var magic PS3

	args := Args{17, 10}
	var reply1 int
	var reply2 Answer // because type is different

	magic.Multiply(&args, &reply1)

	fmt.Println(reply1)

	magic.Divide(&args, &reply2)

	fmt.Println(reply2)

}
