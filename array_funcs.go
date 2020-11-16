package main

// Given a fixed length array arr of integers, duplicate each
// occurrence of zero, shifting the remaining elements to the right.
func duplicateZeros(arr []int) {
	skip := false
	for i := 0; i < len(arr); i++ {
		if skip == false {
			if arr[i] == 0 && (i+1) < len(arr) {
				arr = append(arr[:i+1], append([]int{0}, arr[i+1:len(arr)-1]...)...)
				skip = true
			}
		} else {
			skip = false
		}
	}
}
