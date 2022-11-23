package main

import "fmt"

func main() {
	var n int
	a := 1
	fmt.Print("enter a number: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		a += a * i

	}
	fmt.Print(a)

}
