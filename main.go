package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Slink struct {
	head *Node
	tail *Node
}

// type slinkh struct {
// 	tail *Node
// }

// func (l *slinkh) creat(n int) {
// 	newNode := &Node{data: n}
// 	if l.tail == nil {
// 		l.tail = newNode
// 		l.tail.next = newNode
// 	} else {
// 		newNode.next = l.tail.next
// 		l.tail.next = newNode
// 		l.tail = newNode
// 	}

// }

func (l *Slink) create(n int) {
	newNode := &Node{data: n}
	if l.head == nil {
		l.head = newNode
	} else {
		l.tail.next = newNode
	}
	l.tail = newNode
	l.tail.next = l.head
}

func (l *Slink) display() {
	// temp := l.head
	// for temp != l.tail {
	// 	fmt.Print(temp.data, " ")
	// 	temp = temp.next
	// }
	// fmt.Println(temp.data)

	temp := l.head
	fmt.Println(l.head.data)
	for temp != l.tail {
		fmt.Print(temp.data, "-->")
		temp = temp.next
	}
	fmt.Println(temp.data)
	fmt.Println(l.tail.data)
}

func (l *Slink) len() int {
	temp := l.head
	count := 1
	for temp != l.tail {
		count += 1
		temp = temp.next
	}
	return count
}

// func (l *slinkh) disp() {
// 	p := l.tail
// 	temp := l.tail
// 	for l.tail != temp.next {
// 		temp = temp.next
// 		fmt.Print(temp.data, " ")
// 	}
// 	fmt.Print(p.data)

// }

func (l *Slink) push(n int, pos int) {
	if pos == 0 {
		newNode := &Node{data: n}
		newNode.next = l.head
		l.head = newNode
		l.tail.next = newNode
	} else if pos >= l.len() {
		newNode := &Node{data: n}
		//temp = l.tail
		l.tail.next = newNode
		newNode.next = l.head
		l.tail = newNode
		//fmt.Println(newNode.next.data)
	} else if pos > 0 && pos < l.len() {
		newNode := &Node{data: n}
		p := l.head
		i := 0
		for i < pos-1 {
			p = p.next
			i += 1
		}
		newNode.next = p.next
		p.next = newNode
	}
}
func (l *Slink) pop(pos int) {
	if pos == 0 {
		temp := l.head
		l.head = l.head.next
		l.tail.next = l.head
		temp.next = nil
	} else if pos == l.len()-1 {
		temp := l.head
		for temp.next.next != l.head {
			temp = temp.next
		}
		temp2 := temp
		temp.next = l.head
		l.tail = temp2
	} else if pos > 0 && pos < l.len() {
		temp := l.head
		i := 0
		for i < pos-1 {
			temp = temp.next
			i++
		}
		free := temp.next
		temp.next = temp.next.next
		free.next = nil
	}
}

func (l *Slink) reverse() {
	var prev *Node
	currentNode := l.head
	nextNode := currentNode.next
	for currentNode != l.tail {
		prev = currentNode
		currentNode = nextNode
		nextNode = currentNode.next
		currentNode.next = prev
	}
	l.head.next = l.tail
	temp := l.head
	l.head = l.tail
	l.tail = temp
	temp1 := l.head
	fmt.Println(l.head.data)
	for temp1 != l.tail {
		fmt.Print(temp1.data, "-->")
		temp1 = temp1.next
	}
	fmt.Println(temp1.data)
	fmt.Println(l.tail.data)
}
func main() {
	// l := slinkh{}
	// l.creat(10)
	// l.creat(20)
	// l.creat(30)
	// l.creat(40)
	// l.creat(50)
	// l.disp()
	l := Slink{}
	l.create(10)
	l.create(20)
	l.create(30)
	l.create(40)
	l.create(50)
	l.create(60)
	l.create(70)
	l.push(4, 6)
	l.pop(8)
	b := l.len()
	fmt.Println(b)
	l.display()
	l.reverse()

}
