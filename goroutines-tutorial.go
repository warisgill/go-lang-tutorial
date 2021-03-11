package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("\n\n> ========== Unit 10 : Gorotuines (Threading), Channels  ========= \n\n")
	/*
	 A goroutine is a lightweight thread. It runs concurrently with the parent thread.
	 Since its very lightweight, an appliation can has thousands of go routines running concrrently.
	*/

	fmt.Println("\n> Simplest example of goroutine:")
	go Echo("world") // create  a new gorotuine (i.e., thread) shared same address space
	Echo("hello")

	/*
	 Go "Channels" are the pipes. They are used to pass messages (e.g., strings, integers, etc) between
	 goroutines.They also help to synchronize concurrent goroutines in super easy way.
	 (The data flows in the direction of the arrow.)


	*/

	// Example
	chan1 := make(chan []string) // creating a channel (slice)
	// chan2 := make(chan []string)

	go EchoWithChannel("world", chan1)
	go EchoWithChannel("hello", chan1)

	m1 := <-chan1 // receving message from pipe
	m2 := <-chan1 // receving message from pipe
	// m1, m2 := <-pipe, <-pipe
	fmt.Println("\n>Channels:", m1, m2)

	//Example
	pipe2 := make(chan int)
	var tempslice []int // append work on nil slices
	tempslice = append(tempslice, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3)
	// fmt.Println(tempslice)
	go ComputeSum(tempslice[0:4], pipe2)
	go ComputeSum(tempslice[4:8], pipe2)
	go ComputeSum(tempslice[8:len(tempslice)], pipe2)

	sum1, sum2, sum3 := <-pipe2, <-pipe2, <-pipe2
	totalsum := sum1 + sum2 + sum3
	fmt.Println("\n>Computing sum with channels:", sum1, sum2, sum3, totalsum)

	// Skipped: Buffered Channels.

	// //Advance: range & close
	// fmt.Println("\n> Range & Close (Channels)")
	// c := make(chan int, 10)
	// go fibonacci(cap(c), c)
	// for i := range c {
	// 	fmt.Println(i)
	// }

	// //Skipped: select
	// // Mutex : Mutual Exclussion in Go.

}

func Echo(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func EchoWithChannel(s string, mychannel chan []string) {
	tempslice := make([]string, 5, 5)
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		//fmt.Println(s)
		tempslice = append(tempslice, s)
	}

	mychannel <- tempslice // sending message into channels

}

func ComputeSum(slicearray []int, channel chan int) {
	s := 0
	for _, v := range slicearray {
		s += v
	}
	channel <- s
}

// func computeSumAndAverage(slicearray []int, channel chan int) {
// 	s := 0
// 	for _, v := range slicearray {
// 		s += v
// 	}
// 	channel <- s                   // sending 1st message
// 	channel <- s / len(slicearray) // sending 2nd message from same go routine

// }

// func fibonacci(n int, c chan int) {
// 	x, y := 0, 1
// 	for i := 0; i < n; i++ {
// 		c <- x
// 		x, y = y, x+y
// 	}
// 	close(c)
// }
