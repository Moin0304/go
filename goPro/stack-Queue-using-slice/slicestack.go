package main

import "fmt"

type slice struct {
	a []int
}

func (s *slice) len() int {
	return len(s.a)

}

func (s *slice) isempty() bool {
	return s.len() == 0

}
func (s *slice) push(n int) {

	s.a = append(s.a, n)
}

func (s *slice) pop() int {
	r := s.a[s.len()-1]
	s.a = s.a[0 : s.len()-1]
	// s.a = s.a[0 : len(s.a)-1]
	return r

}

func (s *slice) top() int {
	return s.a[len(s.a)-1]

}
func (s *slice) display() {
	for _, val := range s.a {
		fmt.Printf("%d->", val)
	}

}
func main() {
	s := slice{}
	s.push(5)
	s.push(10)
	s.push(15)
	fmt.Print(s.a)
	fmt.Println()
	fmt.Println(s.pop())
	fmt.Println(s.top())
	s.display()

}
