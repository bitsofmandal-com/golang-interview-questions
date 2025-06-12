package main

import (
	"fmt"
)

// Question: What will be the output of the following Go program?
// Answer:
// x - [0, 2, 3, 4, 5]
// y - [0, 2, 3]
func main() {
	var x []int
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)

	y := x
	x[0] = 0
	x = append(x, 4)
	x = append(x, 5)

	fmt.Println(x)
	fmt.Println(y)
}
