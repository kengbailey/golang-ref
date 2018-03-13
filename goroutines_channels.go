/*
    Kenneth Bailey
    3/13/18

    All about goroutines and channels in Go

*/

package main

import (
    "fmt"
    "sync"
)


var hold sync.WaitGroup

func analyze(s string) {
    fmt.Printf("%s: %d\n", s, len(s))
    hold.Done()
}

func main() {
    phrase :=[]string{"hello", "is", "it", "ken", "you're", "looking", "for"}
    hold.Add(len(phrase))
    for i := 0; i < len(phrase); i++ {
        go analyze(phrase[i])
    }
    hold.Wait()
}
