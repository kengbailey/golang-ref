// Sorting slices : Sort a slice of integers (or structs) using Go's sort package and a custom comparison function.

package main

import (
	"fmt"
	"slices"
	"sort"
)

// create type that implements the sort.Interface
type MyInt []int

func (a MyInt) Len() int           { return len(a) }
func (a MyInt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a MyInt) Less(i, j int) bool { return a[i] < a[j] }

func main() {

	// using slices.sort()
	unsortedSlice := []int{1, 7, 2, 8, 3, 9, 4, 0, 5, 6}
	fmt.Printf("usorted slice: %v\n", unsortedSlice)
	slices.Sort(unsortedSlice)
	fmt.Printf("sorted slice: %v\n\n", unsortedSlice)

	// using custom sort function
	unsortedSlice2 := []int{1, 7, 2, 8, 3, 4, 9, 0, 5, 6}
	fmt.Printf("usorted slice2: %v\n", unsortedSlice2)
	sort.Sort(MyInt(unsortedSlice2))
	fmt.Printf("sorted slice2: %v\n\n", unsortedSlice2)

	// using closure
	unsortedSlice3 := []int{1, 7, 2, 3, 8, 4, 9, 0, 5, 6}
	fmt.Printf("usorted slice3: %v\n", unsortedSlice3)
	sort.Slice(unsortedSlice3, func(i, j int) bool { // implement Less function as closure
		return unsortedSlice3[i] < unsortedSlice3[j]
	})
	fmt.Printf("sorted slice3: %v\n", unsortedSlice3)
}
