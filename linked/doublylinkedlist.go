package main

import "fmt"

type node struct {
	data int
	next *node
	pre  *node
}
type doubly struct {
	head *node
	tail *node
}

func (d *doubly) addfirst(e int) {
	new := &node{e, nil, nil}
	if d.head == nil {
		d.head = new
		d.tail = new
	} else {
		new.next = d.head
		d.head.pre = new
		d.head = new
	}
}
func (d *doubly) addlast(e int) {
	new := &node{e, nil, nil}
	if d.head == nil {
		d.head = new
		d.tail = new
	} else {
		d.tail.next = new
		new.pre = d.tail
		d.tail = new
	}

}
func (d *doubly) addany(e, pos int) {
	new := &node{e, nil, nil}
	p := d.head
	i := 0
	for i < pos-1 {
		p = p.next
		i++
	}
	new.next = p.next
	p.next.pre = new
	p.next = new
	new.pre = p

}
func (d *doubly) removefirst() {
	if d.head == nil {
		fmt.Println("doubly is empty")
		return
	}
	d.head = d.head.next
	d.head.pre = nil
	if d.head == nil {
		d.tail = nil
	}

}
func (d *doubly) removelast() {
	if d.head == nil {
		fmt.Println("doubly is empty")
		return
	}
	d.tail = d.tail.pre
	d.tail.next = nil

}
func (d *doubly) removeany(pos int) {
	p := d.head
	i := 0
	for i < pos-1 {
		p = p.next
		i++
	}
	p.next = p.next.next
	p.next.pre = p

}

func (d *doubly) display() {
	p := d.head
	for p != nil {
		fmt.Print(p.data, "-->")
		p = p.next

	}

}

func main() {
	doub := doubly{}
	doub.addfirst(10)
	doub.addfirst(20)
	doub.display()
	fmt.Println()
	doub.addlast(30)
	doub.addlast(40)
	doub.display()
	fmt.Println()
	doub.addany(50, 2)
	doub.display()
	fmt.Println()
	doub.removefirst()
	doub.display()
	fmt.Println()
	doub.removelast()
	fmt.Println()
	doub.display()
	fmt.Println()
	doub.removeany(1)
	doub.display()

}
