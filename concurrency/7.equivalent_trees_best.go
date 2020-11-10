package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree using a closure, sending
// all values to channel ch
func Walk(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}
	walker(t)
	close(ch)
}

// Same checks wether two trees a and b
// contain the same values.
func Same(a, b *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(a, ch1)
	go Walk(b, ch2)

	for {
		a1, ok1 := <-ch1
		b2, ok2 := <-ch2

		if a1 != b2 || ok1 != ok2 {
			return false
		}

		if !ok1 { // if chan closed
			break
		}
	}

	return true
}

func main() {
	// consolidated function calls
	fmt.Println("1:1", Same(tree.New(1), tree.New(1)))
	fmt.Println("1:2", Same(tree.New(1), tree.New(2)))
}
