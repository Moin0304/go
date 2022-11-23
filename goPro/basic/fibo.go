package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("enter a number: ")
	fmt.Scan(&n)
	a := 0
	b := 1
	c := 0
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		c = a + b
		a = b
		b = c

	}

}
