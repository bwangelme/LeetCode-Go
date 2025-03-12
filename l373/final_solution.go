package l373

import (
	"github.com/bwangelme/LeetCode-Go/lt"
)

type IndexNode struct {
	Num1Idx int
	Num2Idx int
}

/*
## 题目描述

给定两个以 非递减顺序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。

定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。

请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。

## 结题思路

创建一个最小堆
1. 将 nums1[0:k-1] 的数字和 nums2[0] 放入到堆中
2. 从堆中取出最小结果 node，加入结果中，同时取 nums1[node.Num1Idx] + nums2[node.Num2Idx] 放入堆中
3. 重复步骤2k次之后，可以保证已经遍历了 num1[0:k-1] 和 nums2[0:k-1] 所有数字，此时堆输出的k个数字就是最小数字对

## 复杂度分析

K 表示堆的大小
第一步将 k 个数字放入堆中，时间复杂度是 O(k logK)
第2, 3步 while 循环遍历 k 次，从堆中添加，删除元素，时间复杂度是O(k logK)

因此总的时间复杂度是 O(k logK)
空间复杂度是堆的大小，O(K)
*/
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
