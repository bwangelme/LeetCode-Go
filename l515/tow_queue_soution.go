package l515

import (
	"github.com/bwangelme/LeetCode-Go/lt"
	"math"
)

/*
## 解题思路

此解法和 largestValues 类似，只不过不是使用计数表示当前层的节点，而是使用队列来存储当前层的节点

## 复杂度分析

n == 树的节点数

时间复杂度是 O(n)
空间复杂度最坏情况下是 O(n)
*/
func largestValuesTwoQueue(root *TreeNode) []int {
	var (
		q1 = lt.NewQueue[*TreeNode]()
		q2 = lt.NewQueue[*TreeNode]()
	)
	if root != nil {
		q1.Offer(root)
	}

	var (
		res = make([]int, 0)
		max = math.MinInt
	)

	for q1.Len() > 0 {
		node := *(q1.Poll())
		max = lt.Max(max, node.Val)

		if node.Left != nil {
			q2.Offer(node.Left)
		}

		if node.Right != nil {
			q2.Offer(node.Right)
		}

		if q1.Len() <= 0 {
			q1 = q2
			q2 = lt.NewQueue[*TreeNode]()

			res = append(res, max)
			max = math.MinInt
		}
	}

	return res
}
