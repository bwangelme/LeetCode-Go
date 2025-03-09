package l538

import (
	"github.com/bwangelme/LeetCode-Go/lt"
)

/*
ConvertBST

## 题目描述

给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。

提醒一下，二叉搜索树满足下列约束条件：

- 节点的左子树仅包含键小于节点键的节点。
- 节点的右子树仅包含键大于节点键的节点。
- 左右子树也必须是二叉搜索树。

## 解题思路

以右左根的方式遍历一颗二叉树，按遍历顺序对遍历的节点求和，这样求出的和就是比被遍历节点大的节点的和

## 复杂度分析

- N 表示树中节点的个数
- 时间复杂度O(N)
- 空间复杂度O(1)
*/
func ConvertBST(root *lt.TreeNode) *lt.TreeNode {
	// 以右根左的方式遍历一棵树

	var (
		stack = lt.NewStack[*lt.TreeNode]()
		cur   = root
		sum   = 0
	)

	for cur != nil || !stack.Empty() {
		for cur != nil {
			stack.Push(cur)
			cur = cur.Right
		}

		pCur := stack.Pop()
		cur = *pCur
		sum += cur.Val
		cur.Val = sum
		cur = cur.Left
	}

	return root
}
