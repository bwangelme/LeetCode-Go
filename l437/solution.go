package l437

import "github.com/bwangelme/LeetCode-Go/lt"

/*
## 题目描述

给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

## 解题思路

深度优先遍历树， pathSum 表示根节点到当前节点的路径上所有节点值的和，sumCnt map 中存储了 pathSum 出现的次数

sumCnt[pathSum-targetSum] 表示的就是，在根节点到当前节点的路径中，
是否存在某个节点的 pathSum 等于当前节点和-目标求和值，如果有的话。表示从该节点到当前节点的和符合目标求和值。
sumCnt[pathSum-targetSum] 即算出符合目标求和值的路径的数量

每次向上回溯节点的时候，当前节点不再包含再路径中，因此 sumCnt[pathSum] 也要相应地减少1

## 复杂度分析

深度优先遍历，遍历了树的每一个节点，因此时间复杂度是 O(N), N 表示树中节点个数

空间复杂度就是看 sumCnt 使用的空间
使用节省空间的写法，sumCnt 每次遇到值为0的情况下删掉数组，那么最坏的空间复杂度是 O(logN)，表示树的高度
*/
func pathSum(root *lt.TreeNode, targetSum int) int {
	sumCnt := make(map[int]int, 0)
	sumCnt[0] = 1
	return dfs(root, targetSum, sumCnt, 0)
}

func dfs(root *lt.TreeNode, targetSum int, sumCnt map[int]int, pathSum int) int {
	if root == nil {
		return 0
	}

	pathSum = pathSum + root.Val
	count, _ := sumCnt[pathSum-targetSum]

	cnt, _ := sumCnt[pathSum]
	sumCnt[pathSum] = cnt + 1

	count += dfs(root.Left, targetSum, sumCnt, pathSum)
	count += dfs(root.Right, targetSum, sumCnt, pathSum)

	sumCnt[pathSum] = sumCnt[pathSum] - 1

	return count
}
