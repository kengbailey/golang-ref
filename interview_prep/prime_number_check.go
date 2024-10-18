// Prime number check : Write a function to check if a given number is prime.

package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 { // 0, 1 are not prime numbers
		return false
	}
	for i := 2; i*i <= n; i++ { // check if divisible by multiples of any # less than n
		if n%i == 0 {
			return false
		}
	}
	return true // otherwise it is prime!
}

func main() {
	for i := 2; i < 50; i++ {
		fmt.Println(i, " is prime? ", isPrime(i))
	}

}
