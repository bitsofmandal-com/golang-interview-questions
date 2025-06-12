package main

import (
	"fmt"
)

// Question: What will be the output of the following Go program?
// Answer: 3 2 1
func ___main() {
	ch := make(chan int)
	select {
	default:
		fmt.Println(1)
		close(ch)
	case <-ch:
		fmt.Println(2)
	case ch <- 1:
		fmt.Println(3)

	}
}
