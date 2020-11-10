package main

import "fmt"

// "select" statement lets a goroutine wait on multiple communication operations
// Select blocks until one of its cases can run, then it executes
// Select chooses one at random if multiple are ready
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			fmt.Println("sent data")
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	// communication channels
	c := make(chan int)
	quit := make(chan int)

	// reads from channel until done, then sends quit message
	go func() {
		fmt.Println("start")
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // reciever blocks until channel has data
		}
		quit <- 0
		fmt.Println("end")
	}()

	//
	fibonacci(c, quit)
}
