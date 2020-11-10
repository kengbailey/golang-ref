package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	// ch <- 10 // this will block b/c buffer is full

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch) // this will block b/c buffer is now empty

	// Sends to a buffered channel block only when the buffer is full.
	// Receives block when the buffer is empty.
}
