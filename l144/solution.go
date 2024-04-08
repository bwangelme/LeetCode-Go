package l144

import "github.com/bwangelme/LeetCode-Go/lt"

/*
## 题目

给你二叉树的根节点 root ，返回它节点值的 前序 遍历。

提示：

树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100

进阶：递归算法很简单，你可以通过迭代算法完成吗？

## 解题思路

1. 将树的左子树不断入栈，入栈时是根节点, 输出根节点
2. 出栈时，根节点的左子树都已经遍历完了
3. 将栈顶节点的右子树重复 1 2 步骤

这样实现了 根 左 右的遍历顺序

## 复杂度分析

n == 树的节点个数

时间复杂度，将整个树遍历了一遍，时间复杂度是 O(n)
空间复杂度，栈的深度取决于树的高度，因此空间复杂度是 O(logN), 最坏情况下，树是链表，空间复杂度是 O(N)
*/

func preorderTraversal(root *lt.TreeNode) []int {
	var (
		s   = lt.NewStack[*lt.TreeNode]()
		cur = root
		res = make([]int, 0)
	)

	for cur != nil || !s.Empty() {
		for cur != nil {
			res = append(res, cur.Val)
			s.Push(cur)
			cur = cur.Left
		}

		cur = *s.Pop()
		cur = cur.Right
	}

	return res
}
