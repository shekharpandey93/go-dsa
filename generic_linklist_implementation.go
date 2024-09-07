package main

import "fmt"

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type LinkList[T any] struct {
	head *Node[T]
}

func NewLinkList[T any]() *LinkList[T] {
	return &LinkList[T]{}
}

type List[T any] interface {
	Add(item T)
	Remove() T
	isEmpty() bool
	Print()
	InsertAt(index int, item T) bool
	RemoveAt(index int) bool
}

func (l *LinkList[T]) Add(item T) {
	newNode := &Node[T]{Value: item, Next: l.head}
	l.head = newNode
}

func (l *LinkList[T]) InsertAt(index int, item T) bool {
	if index < 0 {
		return false
	}
	newNode := &Node[T]{Value: item}
	// Insert at the head
	if index == 0 {
		newNode.Next = l.head
		l.head = newNode
		return true
	}
	// Traverse to the node just before the target index
	current := l.head
	for i := 0; current != nil && i < index-1; i++ {
		current = current.Next
	}
	if current.Next == nil {
		return false
	}
	newNode.Next = current.Next
	current.Next = newNode
	return true
}

func (l *LinkList[T]) RemoveAt(index int) bool {
	if index < 0 || l.head == nil {
		//var zero T
		return false
	}
	// Remove Head
	if index == 0 {
		//item := l.head.Value
		l.head = l.head.Next
		return true
	}
	current := l.head
	// Traverse to the node just before the target index
	for i := 0; current != nil && i < index; i++ {
		current = current.Next
	}
	if current.Next == nil {
		return false
	}
	//val := current.Next
	current.Next = current.Next.Next
	return true
}

func (l *LinkList[T]) RemoveLast() T {
	var zero T
	return zero
}

func (l *LinkList[T]) Remove() T {
	if l.head == nil {
		var zero T
		return zero
	}
	data := l.head.Value
	l.head = l.head.Next
	return data
}

func (l *LinkList[T]) isEmpty() bool {
	return l.head == nil
}

func (l *LinkList[T]) Print() {
	current := l.head
	for current != nil {
		fmt.Printf("%v -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	list := NewLinkList[int]()

	list.Add(3)
	list.Add(2)
	list.Add(1)

	fmt.Println("Initial list:")
	list.Print()

	// Insert at specific index
	list.InsertAt(1, 4)
	fmt.Println("List after inserting 4 at index 1:")
	list.Print()

	// Remove at specific index
	ok := list.RemoveAt(2)
	if ok {
		fmt.Printf("Removed item at index 2: %v\n")
	}

	fmt.Println("List after removal at index 2:")
	list.Print()

	// Remove from the head
	item := list.Remove()
	fmt.Printf("Removed item from head: %v\n", item)

	fmt.Println("List after removal from head:")
	list.Print()

	fmt.Printf("Is the list empty? %v\n", list.isEmpty())

}
