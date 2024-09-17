package main

import (
	"container/heap"
	"fmt"
)

// Approch we created heap of type integer

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)
	for _, num := range nums {
		fmt.Println("h", h)
		if h.Len() < k {
			heap.Push(h, num)
		} else if num > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, num)
		}
	}
	return (*h)[0]
}

func main() {
	//nums := []int{3, 2, 1, 5, 6, 4}
	//k := 2
	//fmt.Printf("The %dth largest element is: %d\n", k, findKthLargest(nums, k))

	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 2
	fmt.Printf("The %dth largest element is: %d\n", k, findKthLargest(nums, k))
}
