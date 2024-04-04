package l199

/*
## 题目描述

给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

提示:

二叉树的节点个数的范围是 [0,100]
-100 <= Node.val <= 100

## 解题思路

此题和 l513, l515 类似，都是用 BFS 的方式遍历二叉树。

每次遍历完一层后，node 就是当前层的最右节点, 因此将 node 的值添加到结果中即可

## 复杂度分析

n == 树中节点个数

时间复杂度: O(n)
空间复杂度: 最坏是 O(n)
*/

import "github.com/bwangelme/LeetCode-Go/lt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var (
		view = make([]int, 0)
	)
	if root == nil {
		return view
	}

	var (
		q       = lt.NewQueue[*TreeNode]()
		current = 1
		next    = 0
	)
	q.Offer(root)

	for q.Len() > 0 {
		node := *(q.Poll())
		current--

		if node.Left != nil {
			q.Offer(node.Left)
			next++
		}

		if node.Right != nil {
			q.Offer(node.Right)
			next++
		}

		if current == 0 {
			current = next
			next = 0
			view = append(view, node.Val)
		}
	}

	return view
}
