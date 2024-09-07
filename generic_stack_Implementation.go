package main

import "fmt"

type StackGeneric[T any] struct {
	items []T
}

func NewStackGeneric[T any]() *StackGeneric[T] {
	return &StackGeneric[T]{}
}

type StackGenericInf[T any] interface {
	push(item T)
	pop() T
	isEmpty() bool
}

func (s *StackGeneric[T]) push(item T) {
	s.items = append(s.items, item)
}

func (s *StackGeneric[T]) pop() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	}
	lastItem := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return lastItem
}

func (s *StackGeneric[T]) isEmpty() bool {
	return len(s.items) == 0
}

func main() {
	stack := NewStackGeneric[int]()
	stack.push(1)
	stack.push(2)
	stack.push(3)
	fmt.Println(stack)
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	stack.push(4)
	fmt.Println(stack.isEmpty())
}
