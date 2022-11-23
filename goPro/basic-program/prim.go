package main

import "fmt"

func main() {
	var n int
	a := 0
	fmt.Println("enter a number: ")
	fmt.Scanf("%d", n)
	for i := 2; i < n+1; i++ {
		if n%i == 0 {
			a += 1
		}

	}
	fmt.Printf("%d\n", a)

	if a == 1 {
		fmt.Println("its a prime number")
	} else {
		fmt.Println("its not a prime number")
	}

}
