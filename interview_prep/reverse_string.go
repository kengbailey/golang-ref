// Reverse a string : Write a function that takes a string and returns its reverse.
package main

import "fmt"

func main() {
	inputString := "1234567890"
	fmt.Println(reverseString1(inputString))
}

// iterate backwards
func reverseString1(input string) string {
	newString := ""
	for i := len(input) - 1; i >= 0; i-- {
		newString = newString + string(input[i])
	}
	return newString
}
