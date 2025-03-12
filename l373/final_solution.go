package l373

import (
	"github.com/bwangelme/LeetCode-Go/lt"
)

type IndexNode struct {
	Num1Idx int
	Num2Idx int
}

func kSmallestPairsV3(nums1 []int, nums2 []int, k int) [][]int {
	var minHeap = lt.NewHeap[*IndexNode](func(v1 *IndexNode, v2 *IndexNode) bool {
		return nums1[v1.Num1Idx]+nums2[v1.Num2Idx] < nums1[v2.Num1Idx]+nums2[v2.Num2Idx]
	})
	if len(nums2) > 0 {
		for i := 0; i < lt.Min(k, len(nums1)); i++ {
			minHeap.Offer(&IndexNode{Num1Idx: i, Num2Idx: 0})
		}
	}

	res := make([][]int, 0)
	for ; k > 0 && minHeap.Len() > 0; k-- {
		node := minHeap.Poll()
		res = append(res, []int{nums1[node.Num1Idx], nums2[node.Num2Idx]})

		if node.Num2Idx < len(nums2)-1 {
			minHeap.Offer(&IndexNode{Num1Idx: node.Num1Idx, Num2Idx: node.Num2Idx + 1})
		}
	}

	return res
}
