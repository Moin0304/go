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

func (s *stack) append(n int) {
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

func (s *stack) pop() int {
	var a int
	if s.isempty() {
		fmt.Println("the stack is empty")
	}
	a = s.rear.data
	p := s.front
	i := 1
	for i < s.len()-1 {
		p = p.next
		i++

	}
	s.rear = p
	s.rear.next = nil

	s.size--
	return a
}
func (s *stack) top() int {
	if s.isempty() {
		fmt.Println("the stack is empty")
	}
	return s.rear.data
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
	s.append(10)
	s.append(20)
	s.append(30)
	s.display()

	fmt.Print(s.pop())
	fmt.Println()
	s.display()
	fmt.Print(s.top())
	// s.display()

}
