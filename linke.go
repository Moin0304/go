package main

import "fmt"

type node struct {
	data int
	next *node
}

type link struct {
	head *node
	tail *node
}

// func (l *link) push(n int) {
// 	new := &node{n, nil}
// 	new.next = l.head
// 	l.head = new

// }

func (l *link) len() int {
	temp := l.head
	c := 0
	for temp != nil {
		temp = temp.next
		c++
	}
	return c
}
func (l *link) addfirst(n int) {
	new := &node{n, nil}
	if l.head == nil {
		l.head = new
	} else {
		new.next = l.head
		l.head = new

	}

}
func (l *link) display() {
	temp := l.head
	i := 0
	for i < l.len() {
		fmt.Print(temp.data, "-->")
		temp = temp.next
		i++
	}

}

func main() {
	l := link{}
	l.addfirst(10)
	l.addfirst(20)
	l.addfirst(30)
	l.display()

}
