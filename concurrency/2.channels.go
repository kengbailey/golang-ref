package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	fmt.Println(s)
	for _, v := range s {
		sum += v
	}

	c <- sum // send sum to c
}

func main() {
	s := []int{7, 13, 6, 2, 94, 29}

	c := make(chan int)

	// here we sum up to s[:len(s)/2] ... first half
	go sum(s[:len(s)/2], c)
	// here we sum from s[:len(s)/2] to end .. second half
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // recieve sums from c

	fmt.Println(x, y, x+y)

	// Sends and receives block until the other side is ready
	// This allows goroutines to synchronize without explicit locks or condition variables.
}
