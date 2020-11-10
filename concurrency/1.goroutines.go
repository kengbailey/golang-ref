package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("starting ... ")
	go f("alittlebitofthis")
	go f("alittlebitofthat")
	f(fmt.Sprintf("%v", time.Now()))
}
