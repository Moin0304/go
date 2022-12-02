package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (root *Node) PreOrderTraversal() {
	if root != nil {
		fmt.Printf("%d ", root.data)
		root.left.PreOrderTraversal()
		root.right.PreOrderTraversal()
	}
	return
}
func (root *Node) inordertravarsel() {
	if root != nil {
		root.left.inordertravarsel()
		fmt.Printf("%d ", root.data)
		root.right.inordertravarsel()
	}

}
func (root *Node) postorder() {
	if root != nil {
		root.right.postorder()
		root.left.postorder()
		fmt.Printf("%d ", root.data)

	}

}
func main() {
	tree := Node{1, &Node{2, &Node{4, nil, nil}, &Node{5, nil, nil}}, &Node{3, &Node{6, nil, nil}, &Node{7, nil, nil}}}
	fmt.Printf("Pre Order Traversal of the given tree is: ")
	tree.PreOrderTraversal()
	fmt.Println()
	fmt.Printf("In Order Traversal of the given tree is: ")
	tree.inordertravarsel()
	fmt.Println()
	fmt.Printf("post Order Traversal of the given tree is: ")
	tree.postorder()
}
