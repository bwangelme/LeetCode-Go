package l513

/*
## 题目描述

给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。

假设二叉树中至少有一个节点。

提示:

二叉树的节点个数的范围是 [1,104]
-231 <= Node.val <= 231 - 1

## 解题思路

此题和 l515 类似，都是用 BFS 遍历树，每次在层的开始时，记录最左边的值。当遍历完整棵树后，记录的就是最底层，最左边的值

## 复杂度分析

n == 树的节点的个数

时间复杂度: O(n)
空间复杂度: 最坏是 O(n)
*/

import "github.com/bwangelme/LeetCode-Go/lt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	var (
		q          = lt.NewQueue[*TreeNode]()
		current    = 1
		next       = 0
		bottomLeft = root.Val
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
			if q.Len() > 0 {
				bottomLeft = (*q.Peek()).Val
			}
		}
	}

	return bottomLeft
}
