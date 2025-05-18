package l114

import (
	"github.com/bwangelme/LeetCode-Go/lt"
)

/*
## 题目描述

给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。

## 解题思路

1. 使用先序遍历整个树
2. 左子树变成链表返回头和尾
3. 右子树变成链表返回头和尾
4. 拼接链表 root -> LHead - > LTail -> RHead -> RTail
5. 返回链表，向上回溯，上一层继续拼接链表

## 复杂度分析

N = 树的叶子节点个数

- 空间复杂度

没有申请新空间, 空间复杂度是递归深度

递归深度是树的高度，因此空间复杂度是 O(logN)

- 时间复杂度

遍历了所有节点，因此时间复杂度是 O(N)
*/
func flatten(root *lt.TreeNode) {
	flattenTree(root)
}

func flattenTree(root *lt.TreeNode) (head, tail *lt.TreeNode) {
	if root == nil {
		return nil, nil
	}
	if root.Left == nil && root.Right == nil {
		return root, root
	}

	LHead, LTail := flattenTree(root.Left)
	RHead, RTail := flattenTree(root.Right)

	root.Left = nil
	if LHead != nil && RHead != nil {
		root.Right = LHead
		LTail.Right = RHead
		tail = RTail
	} else if LHead == nil && RHead != nil {
		root.Right = RHead
		tail = RTail
	} else if LHead != nil && RHead == nil {
		root.Right = LHead
		tail = LTail
	}

	return root, tail
}

/*
## 题目描述

给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。

## 解题思路

使用循环遍历二叉树，若当前节点存在左子树，找到左子树最右侧节点，将当前节点的右子树连接到该节点右侧，
然后把当前节点的左子树移到右子树位置，同时将左子树置为 nil。接着继续处理下一个节点。

## 复杂度分析

N = 树的节点个数
- 空间复杂度
没有申请新空间, 空间复杂度是 O(1)

- 时间复杂度
遍历了所有节点，因此时间复杂度是 O(N)
*/
func flattenLoop(root *lt.TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			// 找到左子树最右侧的节点
			last := curr.Left
			for last.Right != nil {
				last = last.Right
			}
			// 将当前节点的右子树连接到左子树最右侧节点的右侧
			last.Right = curr.Right
			// 将当前节点的左子树移到右子树位置
			curr.Right = curr.Left
			// 将当前节点的左子树置为 nil
			curr.Left = nil
		}
		// 移动到下一个节点
		curr = curr.Right
	}
}
