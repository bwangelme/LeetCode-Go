package l143

/*
## 题目描述

给定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

提示：

链表的长度范围为 [1, 5 * 104]
1 <= node.val <= 1000

## 解题思路

1. 寻找到链表的中间节点，将链表分成两半，如果链表是奇数个，需要确保前半部分多一个节点
2. 将链表的后半部分反转
3. 将两部分链表连接起来

## 复杂度分析

N 表示链表的长度

寻找中间节点: O(N/2)
反转后半部分: O(N/2)
链接两部分: O(N)

因此时间复杂度是 O(N)

因为没有申请其他和 N 有关的空间，空间复杂度是 O(1)

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var (
		prev *ListNode = nil
		cur            = head
		next *ListNode = nil
	)
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

func link(node1 *ListNode, node2 *ListNode, head *ListNode) {
	var prev = head
	for node1 != nil && node2 != nil {
		temp := node1.Next

		prev.Next = node1
		node1.Next = node2

		prev = node2
		node1 = temp
		node2 = node2.Next
	}

	if node1 != nil {
		prev.Next = node1
	}
}

func reorderList(head *ListNode) {
	var middlePrev = getMiddlePrev(head)
	var temp = middlePrev.Next
	middlePrev.Next = nil

	var res = &ListNode{}
	link(head, reverseList(temp), res)
}

func getMiddlePrev(head *ListNode) *ListNode {
	var (
		dummy = &ListNode{
			Val:  0,
			Next: head,
		}
		slow = dummy
		fast = dummy
	)

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}

	return slow
}
