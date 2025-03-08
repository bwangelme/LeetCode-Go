package l897

import "github.com/bwangelme/LeetCode-Go/lt"

/*
## 题目描述

给你一棵二叉搜索树的 root ，请你 按中序遍历 将其重新排列为一棵递增顺序搜索树，使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右子节点。

## 结题思路

首先将左子树变成一个链表，返回 head, tail

左子树 tail.Right = root
将 root.Left = nil，去除所有节点的 left

然后将右子树变成一个链表, 返回 head, tail
root.Right = 右子树head

## 复杂度分析

- N表示树的节点个数

时间复杂度 O(N)
空间复杂度，栈的递归深度为树的高度，为 O(logN)
*/
func increasingBST(root *lt.TreeNode) *lt.TreeNode {
	head, _ := flattenTree(root)
	return head
}

func flattenTree(root *lt.TreeNode) (head, tail *lt.TreeNode) {
	if root == nil {
		return nil, nil
	}

	if root.Left != nil {
		leftHead, leftTail := flattenTree(root.Left)
		leftTail.Right = root
		head = leftHead
		root.Left = nil
	} else {
		head = root
	}

	if root.Right != nil {
		rightHead, rightTail := flattenTree(root.Right)
		root.Right = rightHead
		tail = rightTail
	} else {
		tail = root
	}

	return head, tail
}
