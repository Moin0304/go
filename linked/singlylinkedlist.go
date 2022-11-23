package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type SinglyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (sl *SinglyLinkedList) len() int {
	return sl.size
}

func (sl *SinglyLinkedList) addlast(n int) {
	new := &Node{n, nil}
	if sl.head == nil {
		sl.head = new
	} else {
		sl.tail.next = new
	}
	sl.tail = new
	sl.size++
}

func (sl *SinglyLinkedList) addfirst(n int) {
	new := &Node{n, nil}
	if sl.head == nil {
		sl.head = new
		sl.tail = new
	} else {
		new.next = sl.head
		sl.head = new
	}
	sl.size++
}

func (sl *SinglyLinkedList) addany(n int, index int) {
	new := &Node{n, nil}
	p := sl.head
	count := 1
	for count < index {
		p = p.next
		count++
	}
	new.next = p.next
	p.next = new
	sl.size++
}

func (sl *SinglyLinkedList) removefirst() {
	if sl.head == nil {
		fmt.Println("Linked list is empty")
	} else {
		p := sl.head.next
		sl.head = p
		sl.size--
	}
	if sl.head == nil {
		sl.tail = nil
	}
}

func (sl *SinglyLinkedList) removelast() {
	if sl.head == nil {
		fmt.Println("Linked list is empty")
	} else {
		p := sl.head
		e := 1
		for e < (sl.len())-1 {
			p = p.next
			e++
		}
		sl.tail = p
		sl.tail.next = nil
		sl.size--
	}
}

func (sl *SinglyLinkedList) removeany(i int) {

	if i != 0 && i != (sl.len())-1 && i < sl.len() {
		if sl.head == nil {
			fmt.Println("Linked list is empty")
		} else {
			p := sl.head
			count := 1
			for count < i {
				p = p.next
				count++
			}
			p.next = p.next.next
			sl.size--
		}
	} else if i == 0 {
		sl.removefirst()
	} else if i == (sl.len())-1 {
		sl.removelast()
	} else {
		fmt.Println("Index ot of range...!")
	}

}

func (sl *SinglyLinkedList) search(key int) int {
	p := sl.head
	index := 0
	for p != nil {
		if p.data == key {
			return index
		}
		p = p.next
		index++
	}
	return -1
}

func (sl *SinglyLinkedList) display() {
	p := sl.head
	for p != nil {
		fmt.Print(p.data, "-->")
		p = p.next
	}
}

func main() {
	l := &SinglyLinkedList{}
	l.addlast(10)
	l.addlast(20)
	l.addlast(30)
	l.display()
	fmt.Println()
	l.removeany(20)
	l.display()
	fmt.Println()
	fmt.Println(l.len())

}
