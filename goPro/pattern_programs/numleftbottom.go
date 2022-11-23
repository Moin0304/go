package main

import "fmt"

func main() {
	var n int
	fmt.Print("enter a number: ")
	fmt.Scan(&n)
	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			if i >= j {
				fmt.Printf("%d", j)
			} else {
				fmt.Print(" ", " ")
			}

		}
		fmt.Println()

	}

}
