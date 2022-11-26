package main

import "fmt"

type queue struct {
	a []int
}

func (q *queue) len() int {
	return q.len()

}

func (q *queue) enque(n int) {
	q.a = append(q.a, n)

}

func (q *queue) deque() int {
	r := q.a[0]
	q.a = q.a[1:]
	return r
}

func (q *queue) first() int {
	return q.a[0]

}
func (q *queue) display() {
	for _, val := range q.a {
		fmt.Printf("%d->", val)
	}
}

func main() {
	q := queue{}
	q.enque(10)
	q.enque(20)
	q.enque(30)
	q.display()
	fmt.Println()
	fmt.Println(q.deque())
	fmt.Println(q.first())
	q.display()
}
