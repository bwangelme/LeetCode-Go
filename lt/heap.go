package lt

import "container/heap"

type Heap[T any] struct {
	queue       []T
	CompareFunc func(v1 T, v2 T) bool
}

func NewHeap[T any](compareFunc func(v1 T, v2 T) bool) *Heap[T] {
	h := &Heap[T]{
		CompareFunc: compareFunc,
	}
	heap.Init(h)
	return h
}

func (h *Heap[T]) Len() int {
	return len(h.queue)
}

func (h *Heap[T]) Less(i, j int) bool {
	v1 := h.queue[i]
	v2 := h.queue[j]
	return h.CompareFunc(v1, v2)
}

func (h *Heap[T]) Swap(i, j int) {
	h.queue[i], h.queue[j] = h.queue[j], h.queue[i]
}

func (h *Heap[T]) Push(val any) {
	h.queue = append(h.queue, val.(T))
}

func (h *Heap[T]) Pop() any {
	if h.Len() <= 0 {
		return nil
	}

	v := h.queue[h.Len()-1]
	h.queue = (h.queue)[:h.Len()-1]
	return v
}

func (h *Heap[T]) Offer(val T) {
	heap.Push(h, val)
}

func (h *Heap[T]) Poll() (res T) {
	val := heap.Pop(h)
	if val != nil {
		res = val.(T)
	}
	return res
}

func (h *Heap[T]) Peek() (res T) {
	if h.Len() > 0 {
		res = h.queue[0]
	}

	return res
}
