package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func Insert(v int, n *Node) *Node {
	if n == nil {
		return &Node{value: v}
	}
	if n.value > v {
		n.left = Insert(v, n.left)
	} else {
		n.right = Insert(v, n.right)
	}
	return n
}

func InOrder(node *Node) {
	if node == nil {
		return
	}
	InOrder(node.left)
	fmt.Println(node.value)
	InOrder(node.right)
}

func PreOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.value)
	PreOrder(node.left)
	PreOrder(node.right)
}

func PostOrder(node *Node) {
	if node == nil {
		return
	}
	PostOrder(node.left)
	PostOrder(node.right)
	fmt.Println(node.value)
}

func Search(node *Node, target int) bool {
	if node == nil {
		return false
	}

	if node.value == target {
		return true
	}

	if target < node.value {
		return Search(node.left, target)
	}

	return Search(node.right, target)

}

func main() {
	var root *Node

	treeInputs := []int{5, 9, 4, 2, 4, 6, 5, 1, 7}

	for _, v := range treeInputs {
		root = Insert(v, root)
	}

	InOrder(root)
	fmt.Println("PreOrder")
	PreOrder(root)
	fmt.Println("PostOrder")
	PostOrder(root)
	fmt.Println("is 6 exist: ", Search(root, 6))
}
