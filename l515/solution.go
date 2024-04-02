package l515

/*
## 题目描述

给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。

提示：

二叉树的节点个数的范围是 [0,104]
-2^31 <= Node.val <= 2^31 - 1

## 解题思路

树要广度优先遍历，bfs, 还是需要用栈来遍历

使用 current 记录当前层的节点数，当 current == 0 的时候，当前层遍历完，输出一个最大值

## 复杂度分析

n == 树的节点个数

时间复杂度: O(n), 需要遍历完树的节点

空间复杂度: 最坏 O(n), 树只有一层，所有节点都放到了队列中
*/

import (
	"github.com/bwangelme/LeetCode-Go/lt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	var (
		current = 0
		next    = 0
		q       = lt.NewQueue[*TreeNode]()
	)

	if root != nil {
		current += 1
		q.Offer(root)
	}

	var (
		res = make([]int, 0)
		max = math.MinInt
	)

	for q.Len() > 0 {
		node := *q.Poll()
		current--
		max = lt.Max(max, node.Val)

		if node.Left != nil {
			q.Offer(node.Left)
			next++
		}

		if node.Right != nil {
			q.Offer(node.Right)
			next++
		}

		if current == 0 {
			res = append(res, max)
			current = next
			next = 0
			max = math.MinInt
		}
	}

	return res
}
