package main

import "fmt"

type node struct {
	pre  *node
	data int
	next *node
}

type doubly struct {
	head *node
	tail *node
	size int
}

func (d *doubly) len() int {
	return d.size

}

func (d *doubly) isempty() bool {
	return d.size == 0

}

func (d *doubly) adf(n int) {
	new := &node{nil, n, nil}
	if d.head == nil {
		d.head = new
		d.tail = new
	} else {
		new.next = d.head
		d.head = new
		new.pre = nil
	}
	d.size++

}

func (d *doubly) adl(n int) {
	new := &node{nil, n, nil}
	if d.isempty() {
		d.head = new
		d.tail = new
	} else {
		d.tail.next = new
		new.pre = d.tail
		d.tail = new
		new.next = nil
	}
	d.size++

}

func (d *doubly) ada(n, i int) {
	new := &node{nil, n, nil}
	if i == 0 {
		d.adf(n)
	} else if i == d.len()-1 {
		d.adl(n)
	} else if i >= d.len() {
		fmt.Println("index is outof range")
	} else {
		p := d.head
		e := 1
		for e < i {
			p = p.next
			e++
		}
		new.next = p.next
		p.next = new
		new.pre = p
		d.size++

	}

}

func (d *doubly) rmf() int {
	var a int
	if d.isempty() {
		fmt.Println("doubly linkedlist is empty")
	} else {
		a = d.head.data
		d.head = d.head.next
		d.size--
	}
	return a

}

func (d *doubly) rml() int {
	var a int
	if d.isempty() {
		fmt.Println("doubly linkedlist is empty")
	} else {
		a = d.tail.data
		d.tail = d.tail.pre
		d.tail.next = nil
		d.size--
	}
	return a

}

func (d *doubly) rma(i int) int {
	var a int
	if i == 0 {
		a = d.rmf()
	} else if i == d.len()-1 {
		a = d.rml()
	} else {
		p := d.head
		e := 0
		for e < i-1 {
			p = p.next
			e++
		}
		a = p.next.data
		p.next = p.next.next
		p.next.pre = p
	}
	return a

}

func (d *doubly) display() {
	p := d.head
	for p != nil {
		fmt.Print(p.data, "-->")
		p = p.next
	}
	fmt.Println()

}

func main() {
	d := doubly{}
	d.adf(10)
	d.adf(20)
	d.adf(30)
	d.display()
	d.adl(20)
	d.adl(30)
	d.display()
	d.ada(50, d.len()-1)
	d.display()
	fmt.Println(d.rmf())
	d.display()
	fmt.Println(d.rml())
	d.display()
	fmt.Println(d.rma(d.len() - 1))
	d.display()

}
