/*
   	Ken Bailey
	2018-11-20

   	All about slices and arrays in Go

*/

package main

import (
	"fmt"
)

func main() {

	// Declare a Slice
	var stringList []string

	// Declare an Array
	// Same as slice but with defined size!
	// Ints initialized to 0
	// Strings initialized to empty strings
	var intArray [4]int
	stringArray := [3]string{"Ken", "was", "here..."}
	otherStringArray := [...]string{"Here", "as", "well..."}

	// Append to slice
	stringList = append(stringList, "Ken")
	stringList = append(stringList, "was", "here...")

	// Append to array
	intArray[0] = 1

	// Print
	fmt.Println(stringList)
	fmt.Println(intArray)
	fmt.Println(stringArray)
	fmt.Println(otherStringArray)
}
