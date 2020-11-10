package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, y+x
	}
	close(c) // only sender should close a channel
	// sending on a closed channel will cause a panic
}

func main() {
	c := make(chan int, 10)

	c <- 123

	// we can test to see if a channel is open by accepting a second argument
	val, okk := <-c
	if !okk {
		fmt.Println("c chan is closed")
	} else {
		fmt.Printf("c chan is open; %v\n", val)
	}

	// here we insert various fib values into our channel
	// cap() returns the numerical value of the capacity of the given slice; so 10 in this case
	go fibonacci(cap(c), c)

	// here we use range to cycle through all values in the channel
	// when it exits the channel is empty, closed and blocking
	for i := range c {
		fmt.Println(i)
	}

	// we can test to see if a channel is open by accepting a second argument
	val, okk = <-c
	if !okk {
		fmt.Println("c chan is closed")
	} else {
		fmt.Printf("c chan is open; %v\n", val)
	}
}
