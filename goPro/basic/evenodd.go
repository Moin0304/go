package main

import "fmt"

func main() {
	var n int
	fmt.Print("enter a number: ")
	fmt.Scan(&n)
	if n%2 == 0 {
		fmt.Print("its even")
	} else {
		fmt.Print("its odd")
	}

}
