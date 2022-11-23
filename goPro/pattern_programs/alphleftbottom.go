package main

import "fmt"

func main() {
	for i := 'A'; i <= 'Z'; i++ {
		for j := 'A'; j <= i; j++ {
			fmt.Printf("%c ", j)
		}
		fmt.Println()
	}
}
