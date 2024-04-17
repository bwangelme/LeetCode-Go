package l814

/*
## 题目描述
814. 二叉树剪枝

给你二叉树的根结点 root ，此外树的每个结点的值要么是 0 ，要么是 1 。

返回移除了所有不包含 1 的子树的原二叉树。

节点 node 的子树为 node 本身加上所有 node 的后代

## 解题思路

对树进行后序遍历

先判断左右子树是否可以剪枝
如果左右子树都可以剪，且当前节点值是0
那么当前节点可以剪枝

## 复杂度分析

N == 树的节点数

### 时间复杂度

因为遍历了每个节点，因此时间复杂度是 O(N)

### 空间复杂度

递归深度为树的高度
最好为 O(logN) ，满二叉树
最坏为 O(N), 树是一个链表
*/

import "github.com/bwangelme/LeetCode-Go/lt"

func pruneTree(root *lt.TreeNode) *lt.TreeNode {
	if root == nil {
		return nil
	}

	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)

	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}

	return root
}
