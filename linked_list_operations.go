package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func main() {
	head := insertNode(1, insertNode(9, insertNode(9, insertNode(9, nil))))
	head = addOne(head)
	//head = insertNNode(head, 3, 10)
	//head = insertNNode(head, 2, 15)
	//head = deleteNode(head, 1)
	//head = reverse(head)
	//fmt.Printf("\n======Reverse=====\n")
	fmt.Printf("%d ", head.data)
	//for head != nil {
	//fmt.Printf("%d ",head.data)
	//head = head.next
	//}

}

func insertNNode(head *Node, index, value int) *Node {
	newNode := &Node{
		data: value,
		next: nil,
	}
	if head.next == nil {
		head.next = newNode
	}
	currentNode := getNNode(head, index)
	holdindingData := currentNode.next
	currentNode.next = newNode
	newNode.next = holdindingData

	return head
}

func getNNode(head *Node, index int) *Node {
	counter := 0
	currentNode := head
	for counter != index-1 {
		currentNode = currentNode.next
		counter++
	}
	return currentNode
}

func insertNode(value int, next *Node) *Node {
	var n Node
	n.data = value
	n.next = next
	return &n
}

func deleteNode(head *Node, index int) *Node {
	currentNode := getNNode(head, index)
	holdindingData := currentNode.next
	currentNode.next = holdindingData.next
	return head
}

func reverse(head *Node) *Node {
	curr := head
	var temp *Node = nil
	var pre *Node = nil

	for curr != nil {
		fmt.Printf("%d ", curr.data)
		temp = curr.next
		curr.next = pre
		pre = curr
		curr = temp
	}
	return pre
}

func addWithCarry(head *Node) int {

	// If linked list is empty, then
	// return carry
	if head == nil {
		return 1
	}

	// Add carry returned be next node call
	res := head.data + addWithCarry(head.next)
	// Update data and return new carry
	head.data = (res) % 10
	return (res) / 10
}

func addOne(head *Node) *Node {
	// Add 1 to linked list from end to beginning
	carry := addWithCarry(head)
	fmt.Println("carry", carry)
	// If there is carry after processing all nodes,
	// then we need to add a new node to linked list
	if carry > 0 {

		newNode := insertNode(carry, nil)
		newNode.next = head
		return newNode // New node becomes head now
	}
	return head
}
