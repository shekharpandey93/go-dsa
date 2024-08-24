package main

import (
	"errors"
	"fmt"
)

// Stack represents a generic stack data structure
type Stack[T any] struct {
	items []T
}

// NewStack creates and returns a new empty stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

// Push adds an item to the top of the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, errors.New("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// Peek returns the top item from the stack without removing it
func (s *Stack[T]) Peek() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty returns true if the stack is empty, false otherwise
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack
func (s *Stack[T]) Size() int {
	return len(s.items)
}

func main() {
	// Example usage with integers
	intStack := NewStack[int]()
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Println("Int Stack size:", intStack.Size())
	if val, err := intStack.Pop(); err == nil {
		fmt.Println("Popped:", val)
	}
	if val, err := intStack.Peek(); err == nil {
		fmt.Println("Peeked:", val)
	}

	// Example usage with strings
	stringStack := NewStack[string]()
	stringStack.Push("Hello")
	stringStack.Push("World")

	fmt.Println("String Stack size:", stringStack.Size())
	if val, err := stringStack.Pop(); err == nil {
		fmt.Println("Popped:", val)
	}
	if val, err := stringStack.Peek(); err == nil {
		fmt.Println("Peeked:", val)
	}
}
