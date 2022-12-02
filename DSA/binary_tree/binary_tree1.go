package main

import "fmt"

type tree struct {
	root *node
}

type node struct {
	element int
	left    *node
	right   *node
}

func (t *tree) insert(data int) {
	if t.root == nil {
		t.root = &node{element: data}
	} else {
		t.root.insert(data)
	}
}

func (n *node) insert(data int) {
	if data <= n.element {
		if n.left == nil {
			n.left = &node{element: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &node{element: data}
		} else {
			n.right.insert(data)
		}
	}
}
func PreOrder(n *node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d ", n.element)
		PreOrder(n.left)
		PreOrder(n.right)
	}
}

func Inorder(n *node) {
	if n == nil {
		return
	} else {
		Inorder(n.left)
		fmt.Printf("%d ", n.element)
		Inorder(n.right)

	}

}
func postorder(n *node) {
	if n == nil {
		return
	} else {
		postorder(n.right)
		postorder(n.left)
		fmt.Printf("%d ", n.element)
	}

}
func main() {
	var t tree

	t.insert(10)
	t.insert(20)
	t.insert(30)
	t.insert(40)
	t.insert(50)
	fmt.Printf("preoder traversal oreder are: ")
	PreOrder(t.root)
	fmt.Println()
	fmt.Printf("Inoder traversal oreder are: ")
	Inorder(t.root)
	fmt.Println()
	fmt.Printf("postoder traversal oreder are: ")
	postorder(t.root)
	fmt.Println()

}
