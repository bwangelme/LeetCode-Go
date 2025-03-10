package l703

import (
	"container/heap"
)

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}

func (h *IntHeap) Peek() int {
	return (*h)[0]
}

func (h *IntHeap) Size() int {
	return len(*h)
}

func NewIntHeap() *IntHeap {
	h := &IntHeap{}
	heap.Init(h)
	return h
}

type KthLargest struct {
	minHeap *IntHeap
	size    int
}

func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{
		minHeap: NewIntHeap(),
		size:    k,
	}
	for _, num := range nums {
		kth.Add(num)
	}
	return kth
}

func (this *KthLargest) Add(val int) int {
	if this.minHeap.Size() < this.size {
		heap.Push(this.minHeap, val)
	} else if val > this.minHeap.Peek() {
		heap.Pop(this.minHeap)
		heap.Push(this.minHeap, val)
	}

	return this.minHeap.Peek()
}
