package main

type Heap struct {
	arr []int
}

func (m *Heap)buildHeap() {

}
func (m *Heap) maxHeapify(size, idx int) {

}

func (mh *Heap) sort() {

}

func newHeap(heapArr []int) *Heap {
return &Heap{}
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}
func parent(i int) int {
	return (i - 1) / 2
}

func main()  {

}
