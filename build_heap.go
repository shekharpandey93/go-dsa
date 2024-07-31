package main

import "fmt"

type MinHeap struct {
	arr []int
}

func (m *MinHeap)buildHeap()  {
	cIdx := (len(m.arr) - 2) / 2
	for idx := cIdx; idx >= 0; idx-- {
		m.heapify(idx)
	}
}
func (m *MinHeap)heapify(idx int) {
	// left idx
	leftIdx := 2*idx + 1
	rightIdx := 2*idx + 2
	minIdx := idx
	if leftIdx < len(m.arr) && m.arr[leftIdx] < m.arr[idx] {
		minIdx = leftIdx
	}

	if rightIdx < len(m.arr) && m.arr[rightIdx] < m.arr[idx] {
		minIdx = rightIdx
	}

	if minIdx == idx {
		return
	}
	m.arr[minIdx], m.arr[idx] = m.arr[idx], m.arr[minIdx]
	m.heapify(minIdx)
}

func newMinHeap(heapArr []int) *MinHeap {
	mh := &MinHeap{
		arr: make([]int, len(heapArr)),
	}
	copy(mh.arr, heapArr)

	return mh
}

func (mh *MinHeap) sort() {
	mh.buildHeap()
	for i := len(mh.arr) - 1; i > 0; i-- {
		mh.arr[i], mh.arr[0] = mh.arr[0], mh.arr[i]
		mh.maxHeapify(i, 0)
	}
}

func main()  {
	arr := []int{6, 5, 3, 2, 8, 10, 9}
	heap := newMinHeap(arr)
	fmt.Println("heap",*heap)
	heap.buildHeap()
	fmt.Println(heap.arr)
}
