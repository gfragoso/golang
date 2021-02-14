package main

import "fmt"

var some = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(some(2, 3))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(2, 3))
}
