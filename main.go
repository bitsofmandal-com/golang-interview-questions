package main

import (
	"fmt"
	"sync"
)

// Question: Write a Go program that uses goroutines and channels to print a sequence of numbers starting from 1 to 100.
// Each goroutine should print whether the number is even or odd. The program should terminate when the number exceeds 100.

// Answer:
func worker(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		num, ok := <-ch
		if !ok {
			return
		}
		if num > 100 {
			close(ch)
			return
		}
		if (num % 2) == 0 {
			fmt.Printf("(even):\t%d\n", num)
		} else {
			fmt.Printf("(odd):\t%d\n", num)
		}
		ch <- num + 1
	}
}

func PrintAllNumberInSequence() {
	var wg sync.WaitGroup
	counter := 1
	ch := make(chan int)

	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go worker(ch, &wg)
	}

	ch <- counter // Start the sequence
	wg.Wait()
}
