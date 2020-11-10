package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree sending all values from
// the tree to the channel.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	return
}

// Same determines wether the trees
// t1 and t2 contain the same values
func Same(t1, t2 *tree.Tree) (same bool) {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go Walk(t1, chan1)
	go Walk(t2, chan2)

	same = true

	for i := 0; i < 10; i++ {
		val1 := <-chan1
		val2 := <-chan2
		if val1 != val2 {
			same = false
		}
	}

	return same
}

func main() {

	newTree1 := tree.New(1)
	newTree2 := tree.New(1)

	fmt.Println(newTree1)
	fmt.Println(newTree2)

	fmt.Println(Same(newTree1, newTree2))

}
