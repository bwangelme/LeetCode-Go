package l129

import "github.com/bwangelme/LeetCode-Go/lt"

/*
## 题目描述

给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
每条从根节点到叶节点的路径都代表一个数字：

例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
计算从根节点到叶节点生成的 所有数字之和 。

叶节点 是指没有子节点的节点。

## 复杂度分析

N == 树中节点数

时间复杂度 O(N)
空间复杂度: O(logN) -- O(N)
*/

func sumNumbers(root *lt.TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *lt.TreeNode, path int) int {
	// 路径数字到了叶子节点才算，如果没到叶子节点，则认为这条路径的数字是0
	if root == nil {
		return 0
	}

	path = path*10 + root.Val
	// 到达叶子节点
	if root.Left == nil && root.Right == nil {
		return path
	}
	return dfs(root.Left, path) + dfs(root.Right, path)
}
