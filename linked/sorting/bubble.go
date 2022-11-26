package main

import "fmt"

func bubble(A []int) {

	n := len(A)
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if A[j] > A[j+1] {
				temp := A[j]
				A[j] = A[j+1]
				A[j+1] = temp
			}
		}
	}
}

func main() {

	A := []int{3, 5, 8, 9, 6, 2}
	fmt.Printf("Before sorting %#v\n", A)
	bubble(A)
	fmt.Printf("After sorting %#v\n", A)

}
