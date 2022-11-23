package main

import "fmt"

type node struct {
	data int
	next *node
}

type cir struct {
	head *node
	tail *node
}

func (c *cir) addfirst(n int) {
	new := &node{n, nil}
	if c.head == nil {

		c.head = new
		c.tail = new
		new.next = new
	} else {
		c.tail.next = new
		new.next = c.head
		c.head = new
	}

}

func (c *cir) addlast(e int) {
	new := &node{e, nil}
	if c.head == nil {
		new.next = new
		c.head = new
	} else {
		new.next = c.tail.next
		c.tail.next = new
	}
	c.tail = new
}
func (c *cir) len() int {
	p := c.head
	count := 0
	for p != nil {
		p = p.next
		count++
	}
	return count
}
func (c *cir) addany(e, pos int) {
	new := &node{e, nil}
	p := c.head
	i := 0
	for i < pos-1 {
		p = p.next
		i++
	}
	new.next = p.next
	p.next = new

}
func (c *cir) removefirst() {
	if c.head == nil {
		fmt.Println("circular linkedlist is empty")
		return
	}
	c.tail.next = c.head.next
	c.head = c.head.next

}
func (c *cir) removelast() {
	if c.head == nil {
		fmt.Println("circular linkedlist is empty")
		return
	}
	p := c.head
	i := 1
	for i < c.len()-1 {
		p = p.next
		i++
	}
	c.tail = p
	p = p.next
	c.tail.next = c.head
}
func (c *cir) removeany(pos int) {
	p := c.head
	i := 0
	for i < pos-1 {
		p = p.next
		i++
	}
	p.next = p.next.next

}

func (c *cir) display() {
	p := c.head
	for p != c.tail {
		fmt.Print(p.data, "-->")
		p = p.next
	}
	fmt.Println()

}
func main() {
	c := cir{}
	c.addfirst(10)
	c.addfirst(20)
	c.addfirst(30)
	c.display()
	fmt.Println()
	c.addlast(40)
	c.addlast(50)
	c.display()
	fmt.Println()
	c.addany(50, 2)
	c.display()
	fmt.Println()
	c.removefirst()
	c.display()
	fmt.Println()
	c.removeany(2)
	c.display()
	fmt.Println()
	b := c.len()
	fmt.Println()
	fmt.Print(b)
	c.removelast()
	c.display()

}
