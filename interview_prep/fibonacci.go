// Fibonacci sequence : Write a function that generates the first n Fibonacci numbers.

package main

import "fmt"

func main() {
	n := 3
	initSlice := []int{}
	for i := 0; i < n; i++ {
		if i == 0 || i == 1 {
			initSlice = append(initSlice, i)
		} else {
			initSlice = append(initSlice, initSlice[len(initSlice)-1]+initSlice[len(initSlice)-2])
		}
	}
	fmt.Printf("%v #s of fib --> %v\n", n, initSlice)
}
