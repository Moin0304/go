package main

import "fmt"

type node struct {
	data int
	next *node
}

type stack struct {
	front *node
	rear  *node
	size  int
}

func (s *stack) len() int {
	return s.size
}

func (s *stack) isempty() bool {
	return s.size == 0
}

func (s *stack) enque(n int) {
	new := &node{n, nil}

	if s.isempty() {
		s.front = new
		s.rear = new
	} else {
		s.rear.next = new
		s.rear = new
	}
	s.size++
}

func (s *stack) deque() int {
	var a int
	if s.isempty() {
		fmt.Println("the stack is empty")
	}
	a = s.front.data
	s.front = s.front.next

	s.size--
	return a
}
func (s *stack) first() int {
	if s.isempty() {
		fmt.Println("the stack is empty")
	}
	return s.front.data
}
func (s *stack) display() {
	p := s.front
	for p != nil {
		fmt.Print(p.data, "->")
		p = p.next
	}
	fmt.Println()

}
func main() {
	s := stack{}
	s.enque(10)
	s.enque(20)
	s.enque(30)
	s.display()

	fmt.Print(s.deque())
	fmt.Println()
	s.display()
	fmt.Print(s.first())
	// s.display()

}
