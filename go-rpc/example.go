package main

import (
	"errors"
	"fmt"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {

	fmt.Printf("Simple Example, \n")
	var magic Arith

	args := Args{17, 10}
	var reply int
	var quot Quotient

	magic.Multiply(&args, &reply)

	fmt.Println(reply)

	magic.Divide(&args, &quot)

	fmt.Println(quot)

}
