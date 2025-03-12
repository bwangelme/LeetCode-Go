package l373

import (
	"fmt"
	"github.com/bwangelme/LeetCode-Go/lt"
)

type Node struct {
	Num1 int
	Num2 int
}

func NodeGt(v1 *Node, v2 *Node) bool {
	return v1.Num1+v1.Num2 > v2.Num1+v2.Num2
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var maxHeap = lt.NewHeap[*Node](NodeGt)
	for i := 0; i < lt.Min(k, len(nums1)); i++ {
		for j := 0; j < lt.Min(k, len(nums2)); j++ {
			loopNode := &Node{Num1: nums1[i], Num2: nums2[j]}
			if maxHeap.Len() < k {
				maxHeap.Offer(loopNode)
			} else {
				topNode := maxHeap.Peek()
				if NodeGt(topNode, loopNode) {
					maxHeap.Poll()
					maxHeap.Offer(loopNode)
				}
			}
		}
	}

	res := make([][]int, 0)
	for _, node := range maxHeap.ToArray() {
		fmt.Println(node)
		res = append(res, []int{node.Num1, node.Num2})
	}

	return res
}
