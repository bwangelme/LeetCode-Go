package l160

import "github.com/bwangelme/LeetCode-Go/lt"

/*
## 题目

给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。

## 结题思路

### 两个链表相交

如果两个链表相交，例如下面的例子

	 a -> b

			-> 1 -> 2 -> 3

x -> y -> z

设:

- 链表 A 在相交之前长度为 a
- 链表 B 在相交之前长度为 b
- 链表 A B 相交的长度为 c

我们定义两个指针，

pA 从 链表A 的头部开始移动，移动到链表尾部后从链表 B 的头部开始移动
pB 从 链表B 的头部开始移动，移动到链表尾部后从链表 A 的头部开始移动

pA 移动 a-1 + c + b + 1 步之后到达相交的节点
pB 移动 b-1 + c + a + 1 步之后到达相交的节点

当 pA == pB 的时候，他们的位置处于相交的节点

### 两个链表不相交
如果两个链表相交，例如下面的例子

	a -> b

x -> y -> z

设:

- 链表 A 的长度为 a
- 链表 B 的长度为 b

我们定义两个指针，

pA 从 链表A 的头部开始移动，移动到链表尾部后从链表 B 的头部开始移动
pB 从 链表B 的头部开始移动，移动到链表尾部后从链表 A 的头部开始移动

pA 移动 a-1 + b + 1 步后到达 B 链表的尾部的下一个节点 null
pB 移动 b-1 + a + 1 步后到达 A 链表的尾部的下一个节点 null

此时 pA == pB == null
*/
func getIntersectionNode(headA, headB *lt.ListNode) *lt.ListNode {
	var (
		pA = headA
		pB = headB
	)
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}

		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}

	return pA
}
