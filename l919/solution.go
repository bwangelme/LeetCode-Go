package l919

/*
## 题目描述

完全二叉树 是每一层（除最后一层外）都是完全填充（即，节点数达到最大）的，并且所有的节点都尽可能地集中在左侧。

设计一种算法，将一个新节点插入到一个完整的二叉树中，并在插入后保持其完整。

实现 CBTInserter 类:

CBTInserter(TreeNode root) 使用头节点为 root 的给定树初始化该数据结构；
CBTInserter.insert(int v)  向树中插入一个值为 Node.val == val的新节点 TreeNode。使树保持完全二叉树的状态，并返回插入节点 TreeNode 的父节点的值；
CBTInserter.get_root() 将返回树的头节点。

提示：

树中节点数量范围为 [1, 1000]
0 <= Node.val <= 5000
root 是完全二叉树
0 <= val <= 5000
每个测试用例最多调用 insert 和 get_root 操作 1e4 次

## 解题思路

将子节点不满的节点放入队列中，每次 Insert 时取队列头部节点进行插入。如果子节点满了，移出队列，将移出节点的子节点继续加入到队列中。

## 复杂度分析

n = 树的节点个数

时间复杂度

Constructor: O(n), 会遍历树中所有节点，将适合节点加入到队列中

Insert: O(1), 只找最近一个节点进行添加动作

Get_Root: O(1)

空间复杂度:

O(n), 这是队列所占用的空间

*/

import (
	"fmt"
	"github.com/bwangelme/LeetCode-Go/lt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	return fmt.Sprintf("%v", t.Val)
}

type CBTInserter struct {
	root *TreeNode
	q    lt.Queue[*TreeNode]
}

func Constructor(root *TreeNode) CBTInserter {
	q := lt.NewQueue[*TreeNode]()
	q.Offer(root)
	for (*q.Peek()).Left != nil && (*q.Peek()).Right != nil {
		node := q.Poll()
		q.Offer((*node).Left)
		q.Offer((*node).Right)
	}

	return CBTInserter{
		q:    q,
		root: root,
	}
}

func (this *CBTInserter) Insert(val int) int {
	parent := *this.q.Peek()
	node := &TreeNode{
		Val: val,
	}
	if parent.Left == nil {
		parent.Left = node
	} else {
		parent.Right = node
		this.q.Poll()
		this.q.Offer(parent.Left)
		this.q.Offer(parent.Right)
	}

	return parent.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
