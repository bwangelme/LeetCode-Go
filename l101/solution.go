package l101

import (
	"strconv"
)

/*
## 题目

给你一个二叉树的根节点 root ， 检查它是否轴对称。

## 复杂度分析

最坏的情况是每个节点都判断一次，所以最坏时间复杂度是 O(N/2) = O(N)

- 空间复杂度

最坏情况下，所有节点都判断一次，共会调用 N/2 次函数, 如果考虑函数的栈空间的话，空间复杂度也是 O(N)
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val string) *TreeNode {
	if val == "null" {
		return nil
	}

	valInt, err := strconv.Atoi(val)
	if err != nil {
		return nil
	}

	return &TreeNode{
		Val:   valInt,
		Left:  nil,
		Right: nil,
	}
}

func isSymmetric(root *TreeNode) bool {
	return isEqual(root.Left, root.Right)
}

func isEqual(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil {
		if left == nil && right == nil {
			return true
		}
		return false
	}

	if left.Val != right.Val {
		return false
	}

	if !isEqual(left.Left, right.Right) {
		return false
	}

	if !isEqual(left.Right, right.Left) {
		return false
	}

	return true
}
