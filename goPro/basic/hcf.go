package main

import (
	"fmt"
)

func main() {
	var a, b, min, n int
	fmt.Print("enter a number: ")
	fmt.Print("enter a number: ")
	fmt.Scan(&a)
	fmt.Scan(&b)
	min := 0
	n := 1
	if a > b {
		min = b
	} else {
		min = a
	}
	for i := 2; i < min+1; i++ {
		if a%i == 0&b%i == 0 {
			n = i
		}

	}
	return n
}
