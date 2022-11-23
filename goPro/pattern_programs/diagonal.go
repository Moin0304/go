package main

import "fmt"

func main() {
	var n int
	fmt.Println("enter a number:")
	fmt.Scan(&n)
	// for i := 0; i < n; i++ {
	// 	for j := 0; j < n; j++ {
	// 		if i == j {
	// 			fmt.Print("*", " ")
	// 		} else {
	// 			fmt.Print(" ", " ")
	// 		}

	// 	}
	// 	fmt.Println()

	// }

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i+j == n-1 {
				fmt.Print("*", " ")
			} else {
				fmt.Print(" ", " ")
			}

		}
		fmt.Println()
	}
}
