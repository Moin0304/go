package main

import "fmt"

type node struct {
	data int
	next *node
}

type cir struct {
	head *node
	tail *node
	size int
}

func (c *cir) len() int {
	return c.size

}

func (c *cir) isempty() bool {
	return c.size == 0

}

func (c *cir) adf(n int) {
	new := &node{n, nil}
	if c.isempty() {
		c.head = new
		c.tail = new
		new.next = c.head
	} else {
		new.next = c.head
		c.head = new
		c.tail.next = c.head

	}
	c.size += 1

}

func (c *cir) adl(n int) {
	new := &node{n, nil}
	if c.isempty() {
		c.head = new
		c.tail = new
		c.tail.next = c.head
	} else {
		c.tail.next = new
		c.tail = new
		c.tail.next = c.head
	}
	c.size += 1

}

func (c *cir) ada(n, i int) {

	if i == 0 {
		c.adf(n)
	} else if i >= c.len() {
		c.adl(n)
	} else {
		new := &node{n, nil}
		p := c.head
		e := 0
		for e < i-1 {
			p = p.next
			e += 1
		}
		new.next = p.next
		p.next = new
		c.size += 1
	}

}

func (c *cir) rmf() int {
	var a int
	if c.isempty() {
		fmt.Println("circular linkedlist is empty")
	} else {
		a = c.head.data
		c.head = c.head.next
		c.tail.next = c.head

		c.size -= 1

	}
	if c.isempty() {
		c.head = nil
		c.tail = nil
	}
	return a

}

func (c *cir) rml() int {
	var a int
	if c.isempty() {
		fmt.Println("circular linkedlist is empty")
	} else {
		p := c.head
		i := 1
		for i < c.len()-1 {
			p = p.next
			i += 1
		}
		a = p.next.data
		c.tail = p
		c.tail.next = c.head
		c.size -= 1

	}
	if c.isempty() {
		c.head = nil
		c.tail = nil
	}
	return a
}

func (c *cir) rma(i int) int {
	var a int
	if i == 0 {
		a = c.rmf()

	} else if i >= c.len() {
		fmt.Println("list index outoff range")
	} else if i == c.len()-1 {
		a = c.rml()
	} else {
		p := c.head
		e := 1
		for e < i {
			p = p.next
			e += 1
		}
		a = p.next.data
		p.next = p.next.next
		c.size -= 1

	}
	if c.isempty() {
		c.head = nil
		c.tail = nil
	}
	return a
}

func (c *cir) display() {
	p := c.head
	i := 0
	for i < c.len() {
		fmt.Print(p.data, "-->")
		p = p.next
		i += 1
	}
	fmt.Println()

}

func main() {
	c := cir{}
	c.adf(10)
	c.adf(20)
	c.adf(30)
	c.display()
	c.adl(20)
	c.adl(30)
	c.display()
	c.ada(50, 7)
	c.display()
	fmt.Println(c.rmf())
	c.display()
	c.adf(10)
	c.display()
	fmt.Println(c.rmf())
	c.display()
	fmt.Println()
	// fmt.Println(c.head.data)
	fmt.Println(c.rml())
	c.display()
	fmt.Println()
	fmt.Println(c.rma(c.len() - 1))
	c.display()
}
