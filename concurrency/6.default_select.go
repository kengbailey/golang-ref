package main

import (
	"fmt"
	"time"
)

func main() {

	// <-chan time
	// "tick" is a channel that is the sent the current time
	tick := time.Tick(100 * time.Millisecond) //

	// <-chan time
	// "tock" is a channel that is sent the current time after
	// 		given time has elapsed
	tock := time.After(500 * time.Millisecond)

	for {
		select {
		case currentTime := <-tick: // using case statement to recieve on channel
			fmt.Println("tick.", currentTime)
		case afterTime := <-tock:
			fmt.Println("TOCK!", afterTime)
			return
		default: // default case selected when others are blocking
			fmt.Println("   .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
