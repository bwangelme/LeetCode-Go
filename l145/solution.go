package l145

/*
## 题目描述

给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。

提示：

树中节点的数目在范围 [0, 100] 内
-100 <= Node.val <= 100


进阶：递归算法很简单，你可以通过迭代算法完成吗？

## 复杂度分析

n == 树中节点个数
h == 树的高度

时间复杂度: O(n)
空间复杂度: O(h), 栈的最大空间是树的高度

二叉树的高度 h 最好情况下是 log_2{n+1}, 最坏情况下是 n

因此 O(logN) <= O(h) <= O(n)
*/

import "github.com/bwangelme/LeetCode-Go/lt"

func postorderTraversal(root *lt.TreeNode) []int {
	var (
		res               = make([]int, 0)
		cur               = root
		prev *lt.TreeNode = nil
		s                 = lt.NewStack[*lt.TreeNode]()
	)

	for cur != nil || !s.Empty() {
		for cur != nil {
			s.Push(cur)
			cur = cur.Left
		}

		cur = *(s.Top())
		if cur.Right != nil && cur.Right != prev {
			// 如果右子树还没遍历过，先去遍历右子树
			cur = cur.Right
		} else {
			// 遍历当前节点
			s.Pop()
			res = append(res, cur.Val)
			prev = cur
			// cur == nil，下次循环时，就会从栈中取节点，这样相当于让 cur 往上跳到了它的副节点
			cur = nil
		}
	}

	return res
}
