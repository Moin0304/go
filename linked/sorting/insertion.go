package main

import "fmt"

func insertion(A []int) {
	n := len(A)
	for i := 1; i < n; i++ {
		cvalue := A[i]
		pos := i
		for pos > 0 && A[pos-1] > cvalue {
			A[pos] = A[pos-1]
			pos = pos - 1
		}
		A[pos] = cvalue
	}
}
func main() {

	A := []int{3, 5, 8, 9, 6, 2}
	fmt.Printf("Befor sorting %#v\n", A)
	insertion(A)
	fmt.Printf("After sorting %#v\n", A)
}
