// Find duplicates in a slice : Write a program to find duplicate elements in a slice of integers.

package main

import (
	"fmt"
	"slices"
)

func main() {
	// O(n log n) + O(n) = O(n log n) ... time complexity b/c we are sorting, then iterating
	// to have O(n) we could use a map!
	mySlice := []int{1, 2, 3, 3, 4, 5, 6, 6, 7, 8, 9, 9, 0}
	slices.Sort(mySlice)
	duplicateIndex := []int{}
	for i := 1; i < len(mySlice); i++ {
		if mySlice[i] == mySlice[i-1] { // duplicate
			duplicateIndex = append(duplicateIndex, i)
		}
	}
	fmt.Println("Slice: ", mySlice)
	for _, v := range duplicateIndex {
		fmt.Printf("duplicate # [%v] @ %v\n", mySlice[v], v)
	}
}
