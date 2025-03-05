package l437

import "github.com/bwangelme/LeetCode-Go/lt"

/*
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
