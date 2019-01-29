/*
Kenneth Bailey
2019-01-02

Golang "Worker Pool" concurrency pattern

https://gobyexample.com/worker-pools
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "finished job", j)
		results <- j * 2
	}
	wg.Done()
}

func main() {
	jobs := make(chan int, 100)    // channel to send work
	results := make(chan int, 100) // channel to recieve results

	// kick off workers
	// worker will block until channel populates
	// waitgroup signals that worker is done
	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// send 5 jobs to the channel
	// close channel to signal goroutine stop
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// wait for jobs to finish
	wg.Wait()

	// read result from results channel
	for a := 1; a <= 5; a++ {
		fmt.Println(<-results)
	}
}
