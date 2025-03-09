package l110

import (
	"github.com/bwangelme/LeetCode-Go/lt"
)

func recur(root *lt.TreeNode) int {
	if root == nil {
		return 0
	}
	left := recur(root.Left)
	if left == -1 {
		return -1
	}
	right := recur(root.Right)
	if right == -1 {
		return -1
	}
	if lt.Abs(left-right) <= 1 {
		return lt.Max(left, right) + 1
	} else {
		return -1
	}
}

func isBalanced(root *lt.TreeNode) bool {
	return recur(root) != -1
}
