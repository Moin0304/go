package main

import "fmt"

func selection(A []int) {
	n := len(A)
	for i := 0; i < n-1; i++ {
		pos := i
		for j := i + 1; j < n; j++ {
			if A[j] < A[pos] {
				pos = j
			}

		}
		temp := A[i]
		A[i] = A[pos]
		A[pos] = temp
		fmt.Println(A)

	}

}

func main() {

	A := []int{3, 5, 8, 9, 6, 2}
	fmt.Printf("Before sorting : %#v\n", A)
	selection(A)
	fmt.Printf("after sorting: %#v", A)

}
