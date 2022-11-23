package main

import "fmt"

func isPalindrome(x int) bool {
	a := x
	rev := 0
	rem := 0

	for a > 0 {
		rem = a % 10
		rev = rev*10 + rem
		a /= 10

	}
	if x == rev {
		return true
	} else {
		return false
	}
}

func main() {
	a := isPalindrome(-121)
	fmt.Println(a)

}
