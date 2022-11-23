// package main

// import "fmt"

// func main() {
// 	var a, b int
// 	fmt.Print("enter a number: ")
// 	fmt.Scan(&a)
// 	fmt.Print("enter 2nd number: ")
// 	fmt.Scan(&b)
// 	fmt.Println("before swapping")
// 	a = a + b
// 	b = a - b
// 	a = a - b
// 	fmt.Println("after swapping")
// 	fmt.Println(a)
// 	fmt.Println(b)

// }

package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("enter a number: ")
	fmt.Scan(&a)
	fmt.Print("enter a 2nd number: ")
	fmt.Scan(&b)
	fmt.Print("enter a 3rd number: ")
	fmt.Scan(&c)
	a = a ^ b ^ c
	b = a ^ b ^ c
	c = a ^ b ^ c
	a = a ^ b ^ c
	fmt.Println(a, b, c)

}
