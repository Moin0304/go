package main

import "fmt"

type node struct {
	data int
	next *node
}

type linked struct {
	head *node
	tail *node
}

func (l *linked) addlast(n int) {
	b := node{n, nil}
	new := &b
	if l.head == nil {
		l.head = new
	} else {
		l.tail.next = new
	}
	l.tail = new

}

func (l *linked) display() {
	d := l.head
	for d != nil {
		fmt.Print(d.data, "-->")
		d = d.next
	}

}

func main() {
	a := linked{}
	a.addlast(10)
	a.addlast(20)
	a.display()

}
