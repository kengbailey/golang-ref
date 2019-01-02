/*
    Kenneth Bailey
    3/13/18

    All about goroutines and channels in Go

    Notes:
    - A goroutine is a lightweight thread of execution
    - Goroutines are multiplexed to fewer number of OS threads
    - There may be only on thread in a program with thousands of goroutines
    - If any goroutine in a thread blocks(ex. waiting for user input), then
        another OS thread is created and the remaining goroutines are moved
        to the new OS thread
    - The main goroutine should be running for any other goroutines to run;
        use a sync.WaitGroup to wait on goroutines
    - Goroutines communicate using channels
    - Channels have a type associated with them; only that type can go through
    - Sends and receives through a channel are blocking by default.
    - Channels, by design, prevent race conditions from happening when accessing
        a shared resource ->> When a data is sent to a channel, the control is 
        blocked in the send statement until some other goroutine reads from that 
        channel. Similarly when data is read from a channel, the read is blocked
        until some goroutine writes data to that channel!!!!!
    - When using channels you must create the CONSUMER before the PRODUCER, read
        from channel before you write to channel

*/

package main

import (
    "fmt"
    "sync"
)

// We need a WaitGroup to wait for the goroutines to finish.
// We declare it outside the main function so all functions have access to it.
var hold sync.WaitGroup

// Function we will be running as a goroutine
func analyze(s string) {
    fmt.Printf("%s: %d\n", s, len(s))
    hold.Done()
}

func main() {
    phrase :=[]string{"hello", "is", "it", "ken", "you're", "looking", "for"}

    // We have to tell the WaitGroup how many goroutines to wait on
    hold.Add(len(phrase))

    // Here we kick off all goroutines using the "go" keyword
    for i := 0; i < len(phrase); i++ {
        go analyze(phrase[i])
    }

    // Wait for all goroutines to finish - if not program exits and 
    //  all goroutines are killed when the main program quits.
    hold.Wait()

    // Channel that only allows strings to pass.
    // Channels can only be created using the make function.
    printChan := make(chan string)

    // Here's an anonymous function that takes a channel as a parameter,
    //  reads from that channel and then prints the string read in.
    // Again, we must read from the channel before we write to the channel.
    // Reading from a channel is blocking, so we do it in a goroutine.
    go func(stringChan chan string) {
        tempString := <- stringChan
        fmt.Println(tempString)
    }(printChan)

    // Write to channel
    printChan <- "I can see clearly now, the rain is gone."
}
