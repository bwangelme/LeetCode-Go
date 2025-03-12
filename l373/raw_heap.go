package l373

import (
	"container/heap"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type NodeHeap []*Node

func (h *NodeHeap) Less(i, j int) bool {
	return (*h)[i].Num1+(*h)[i].Num2 > (*h)[j].Num1+(*h)[j].Num2
}

func (h *NodeHeap) Len() int {
	return len(*h)
}

func (h *NodeHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *NodeHeap) Push(x any) {
	*h = append(*h, x.(*Node))
}

func (h *NodeHeap) Pop() any {
	if h.Len() <= 0 {
		return nil
	}

	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

func (h *NodeHeap) Seek() any {
	if h.Len() <= 0 {
		return nil
	}
	return (*h)[0]
}

func (h *NodeHeap) ToArr() []*Node {
	return *h
}

/*
# 实现思路

此方法中，使用了 golang 原生的堆，没有使用泛型，在 leetcode 中依然遇到了超时的错误
*/
func kSmallestPairsV2(nums1 []int, nums2 []int, k int) [][]int {
	var maxHeap = new(NodeHeap)
	heap.Init(maxHeap)
	for i := 0; i < Min(k, len(nums1)); i++ {
		for j := 0; j < Min(k, len(nums2)); j++ {
			loopNode := &Node{Num1: nums1[i], Num2: nums2[j]}
			if maxHeap.Len() < k {
				heap.Push(maxHeap, loopNode)
			} else {
				topNode := maxHeap.Seek().(*Node)
				if topNode.Num1+topNode.Num2 > loopNode.Num1+loopNode.Num2 {
					heap.Pop(maxHeap)
					heap.Push(maxHeap, loopNode)
				}
			}
		}
	}

	res := make([][]int, 0)
	for _, node := range maxHeap.ToArr() {
		res = append(res, []int{node.Num1, node.Num2})
	}

	return res
}
