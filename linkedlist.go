package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linked struct {
	head *node
	tail *node
}

func (l *linked) push(n int) {
	new := &node{data: n}
	if l.head == nil {
		l.head = new

	} else {
		l.tail.next = new
	}
	l.tail = new
}

func (l *linked) search(e int) {
	p := l.head
	i := 0
	for p != nil {
		if p.data == e {
			return
		}
		p = p.next
		i++
	}
	return

}

func (l *linked) addfirst(e int) {
	new := &node{e, nil}
	if l.head == nil {
		l.head = new
	} else {
		new.next = l.head
	}
	l.head = new

}
func (l *linked) addlast(e int) {
	new := &node{e, nil}
	if l.head == nil {
		l.head = new
	} else {
		l.tail.next = new
	}
	l.tail = new

}
func (l *linked) addany(e, p int) {
	new := &node{e, nil}
	temp := l.head
	i := 0
	for i < p-1 {
		temp = temp.next
		i++
	}
	new.next = temp.next
	temp.next = new
}
func (l *linked) removefirst() int {
	var e int
	if l.head == nil {
		fmt.Print("list is empty")

	} else {
		e = l.head.data
		l.head = l.head.next

	}
	return e
}
func (l *linked) len() int {
	temp := l.head
	c := 0
	for temp != nil {
		temp = temp.next
		c++
	}
	return c

}

func (l *linked) removelast() {
	if l.head == nil {
		fmt.Print("list is empty")
		return
	}
	temp := l.head
	i := 1
	for i < l.len()-1 {
		temp = temp.next
		i++
	}
	l.tail = temp
	temp = temp.next
	l.head.data = l.head.data
	l.tail.next = nil
	return
}
func (l *linked) removeany(p int) {
	if l.head == nil {
		fmt.Print("list is empty")
		return
	}
	temp := l.head
	i := 0
	for i < p {
		temp = temp.next
		i++
	}
	temp.data = temp.next.data
	temp.next = temp.next.next
	return

}
func (l *linked) display() {
	p := l.head
	for p != nil {
		fmt.Print(p.data, "-->")
		p = p.next
	}

}
func main() {
	l := linked{}
	l.push(10)
	l.push(20)
	l.push(30)
	l.addfirst(40)
	l.addfirst(50)
	l.addlast(30)
	l.addlast(50)
	l.addany(90, 2)
	l.display()
	fmt.Println()
	fmt.Println(l.removefirst())
	l.display()
	fmt.Println()
	b := l.len()
	fmt.Println(b)
	l.removelast()
	l.display()
	fmt.Println()
	l.removeany(2)
	l.display()

}
