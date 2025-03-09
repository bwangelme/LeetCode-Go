package l653

import "github.com/bwangelme/LeetCode-Go/lt"

type BSTReverseIterator struct {
	stack lt.Stack[*lt.TreeNode]
	cur   *lt.TreeNode
}

func NewBSTReverseIterator(root *lt.TreeNode) BSTReverseIterator {
	return BSTReverseIterator{
		stack: lt.NewStack[*lt.TreeNode](),
		cur:   root,
	}
}

func (this *BSTReverseIterator) Next() int {
	for this.cur != nil {
		this.stack.Push(this.cur)
		this.cur = this.cur.Right
	}

	pCur := this.stack.Pop()
	cur := *pCur
	this.cur = cur.Left
	return cur.Val
}

func (this *BSTReverseIterator) HasNext() bool {
	return !this.stack.Empty() || this.cur != nil
}

type BSTIterator struct {
	stack lt.Stack[*lt.TreeNode]
	cur   *lt.TreeNode
}

func NewBSTIterator(root *lt.TreeNode) BSTIterator {
	return BSTIterator{
		stack: lt.NewStack[*lt.TreeNode](),
		cur:   root,
	}
}

func (this *BSTIterator) Next() int {
	for this.cur != nil {
		this.stack.Push(this.cur)
		this.cur = this.cur.Left
	}

	pCur := this.stack.Pop()
	cur := *pCur
	this.cur = cur.Right
	return cur.Val
}

func (this *BSTIterator) HasNext() bool {
	return !this.stack.Empty() || this.cur != nil
}

/*
## 题目描述

给定一个二叉搜索树 root 和一个目标结果 k，如果二叉搜索树中存在两个元素且它们的和等于给定的目标结果，则返回 true。

二叉搜索树中节点的值都是唯一的

## 结题思路

next 是对二叉搜索树进行 从小到大 遍历的迭代器
prev 是对二叉搜索树进行 从大到小 遍历的迭代器

从树的两端开始查找目标数值，如果大了，就向左移动末尾最大的数，如果小了，就向右移动最小的数

这样就能找到和等于目标值的两个节点

## 复杂度分析

# N 表示树中节点的个数

时间复杂度，只遍历了一次树，因此时间复杂度是 O(N)
空间复杂度，两个迭代器使用的栈的空间是树的高度，因此空间复杂度是 O(logN)
*/
func findTarget(root *lt.TreeNode, k int) bool {
	next := NewBSTIterator(root)
	prev := NewBSTReverseIterator(root)

	nVal := next.Next()
	pVal := prev.Next()
	for nVal != pVal {
		sum := nVal + pVal
		if sum == k {
			return true
		}
		if sum > k {
			pVal = prev.Next()
		} else {
			nVal = next.Next()
		}
	}

	return false
}
