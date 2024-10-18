// Map frequency : Write a program that counts the frequency of words in a string.

package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "The quick brown fox jumps over the lazy dog and the dog runs fast but the fox is very lazy"
	words := strings.Split(text, " ")
	wordCount := make(map[string]int)
	for _, v := range words {
		val, ok := wordCount[v]
		if !ok {
			wordCount[v] = 1
		} else {
			wordCount[v] = val + 1
		}
	}

	for k, v := range wordCount {
		fmt.Println(k, ": ", v)
	}
}
