package main

import "fmt"

type Tree struct {
	data  int
	left  *Tree
	right *Tree
}
func main()  {
	root := &Tree{data: 2}
	root.right = &Tree{data: 7}
	root.left = &Tree{data: 1}
	root.right.right = &Tree{data: 8}
	root.right.right.right = &Tree{data: 9}
	fmt.Println(isBST(root))
}

func isBST(root *Tree) bool {
	if root == nil {
		return true
	}
	if ( (root.left) != nil && (root.left.data > root.data)) {
		return false
	}

	if ((root.right != nil) && (root.right.data < root.data)) {
		return false
	}

	if !isBST(root.left) || !isBST(root.right) {
		return false
	}

	return true
}
