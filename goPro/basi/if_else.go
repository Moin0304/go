package main

import "fmt"

func main() {
	// a := 101
	// if a > 18 && a < 100 {
	// 	fmt.Printf("he is eligible for voting")
	// } else if a == 18 {
	// 	fmt.Printf("he is not eligible for voting,he has to wait for another 6 months")
	// } else {
	// 	fmt.Printf("he is not eligible")
	// }
	n := make([]int, 5, 7)
	var i = 0
	for i < 5 {
		n = append(n, i)
		i++
	}
	fmt.Println(cap(n))

}
