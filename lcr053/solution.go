package lcr053

import "github.com/bwangelme/LeetCode-Go/lt"

/*
## 题目描述

给定一棵二叉搜索树和其中的一个节点 p ，找到该节点在树中的中序后继。如果节点没有中序后继，请返回 null 。

节点 p 的后继是值比 p.val 大的节点中键值最小的节点，即按中序遍历的顺序节点 p 的下一个节点。

## 结题思路

在二叉搜索树中进行搜索，

1. 大于目标标记 res，并向左找，小于目标向右找
2. 找到叶子节点的下一个时，此时 res 标记的就是最小的大于 p 的节点

## 复杂度分析

- N 表示树中节点的个数

时间复杂度 O(logN)
空间复杂度，没有额外空间，O(1)
*/
func inorderSuccessor(root *lt.TreeNode, p *lt.TreeNode) *lt.TreeNode {
	var (
		cur              = root
		res *lt.TreeNode = nil
	)
	for cur != nil {
		if cur.Val > p.Val {
			res = cur
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	return res
}
