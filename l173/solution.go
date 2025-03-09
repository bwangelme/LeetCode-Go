package l173

import "github.com/bwangelme/LeetCode-Go/lt"

/*
BSTIterator

## 题目描述

实现一个二叉搜索树迭代器类BSTIterator ，表示一个按中序遍历二叉搜索树（BST）的迭代器：
- BSTIterator(TreeNode root) 初始化 BSTIterator 类的一个对象。
- BST 的根节点 root 会作为构造函数的一部分给出。指针应初始化为一个不存在于 BST 中的数字，且该数字小于 BST 中的任何元素。
- boolean hasNext() 如果向指针右侧遍历存在数字，则返回 true ；否则返回 false 。
- int next()将指针向右移动，然后返回指针处的数字。

注意，指针初始化为一个不存在于 BST 中的数字，所以对 next() 的首次调用将返回 BST 中的最小元素。

## 结题思路

将中序遍历树的代码变成一个类，外部循环的条件就是 hasNext, 外部循环的操作就是返回一个节点 Next

## 复杂度分析

# N 表示树中节点的个数

空间复杂度，stack 最大为树的高度，因此空间复杂度是 O(logN)

时间复杂度:

hasNext 是 O(1)
Next 每次返回一个树的节点，在遍历过程中调用N次，因此平均时间复杂度是 O(1)
*/
type BSTIterator struct {
	stack lt.Stack[*lt.TreeNode]
	cur   *lt.TreeNode
}

func Constructor(root *lt.TreeNode) BSTIterator {
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
