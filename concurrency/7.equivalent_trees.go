package main

import (
	"fmt"
	"reflect"
	"sort"

	"golang.org/x/tour/tree"
)

// Walk walks the tree sending all values from
// the tree to the channel.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	return
}

// Same determines wether the trees
// t1 and t2 contain the same values
func Same(t1, t2 *tree.Tree) bool {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go Walk(t1, chan1)
	go Walk(t2, chan2)

	out1 := []int{}
	out2 := []int{}

	for i := 0; i < 10; i++ {
		val1 := <-chan1
		out1 = append(out1, val1)
		val2 := <-chan2
		out2 = append(out2, val2)
	}

	sort.Ints(out1)
	sort.Ints(out2)

	return reflect.DeepEqual(out1, out2)
}

func main() {

	newTree1 := tree.New(1)
	newTree2 := tree.New(1)

	fmt.Println(newTree1)
	fmt.Println(newTree2)

	fmt.Println(Same(newTree1, newTree2))

}
