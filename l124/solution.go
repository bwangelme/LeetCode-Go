package l124

import (
	"github.com/bwangelme/LeetCode-Go/lt"
	"math"
)

/*
## 题目描述

二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其 最大路径和 。

## 解题思路

假设一棵树为

	        r
	     /       \
	(l Tree)   (r Tree)

其中 l 和 r 是子树，它可能包括0-N 个节点

定义两个概念

maxSum 以 r 为终点的路径的最大和，如果子树的路径和是负数，放弃计算子树路径
maxLoopSum 不以 r 为终点的路径的最大和，这条路径从 l 的某个节点开始，到 r 的某个节点结束

最终我们求得是整棵树的 maxLoopSum

## 复杂度分析

时间复杂度: O(N)
空间复杂度: 没有额外申请空间，空间复杂度主要源自函数栈的空间，这个最大是树的高度，为 O(logN)
*/
func maxPathSum(root *lt.TreeNode) int {
	_, maxLoopSum := dfs(root, math.MinInt)
	return maxLoopSum
}

func dfs(root *lt.TreeNode, maxLoopSum int) (int, int) {
	if root == nil {
		return 0, maxLoopSum
	}

	leftMaxSum, leftMaxLoopSum := dfs(root.Left, math.MinInt)
	leftMaxSum = lt.Max(0, leftMaxSum)
	rightMaxSum, rightMaxLoopSum := dfs(root.Right, math.MinInt)
	rightMaxSum = lt.Max(0, rightMaxSum)

	maxLoopSumRes := lt.Max(leftMaxLoopSum, rightMaxLoopSum)
	maxLoopSumRes = lt.Max(maxLoopSumRes, root.Val+leftMaxSum+rightMaxSum)

	return root.Val + lt.Max(leftMaxSum, rightMaxSum), maxLoopSumRes
}
