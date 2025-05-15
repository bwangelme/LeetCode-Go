package l24

import (
	"github.com/bwangelme/LeetCode-Go/lt"
)

/*
## 题目描述

给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

输入 1 2 3 4
输出 2 1 4 3

## 解题思路

遍历整个链表，两两进行交换，交换后记录上一次交换的最后一个节点，下次交换后，改变上次交换的最后一个节点的 next

## 复杂度分析

N = 链表长度

空间复杂度，没有申请新空间，O(1)
时间复杂度 O(N)
*/
func swapPairs(head *lt.ListNode) *lt.ListNode {
	var (
		cur               = head
		prev *lt.ListNode = nil
	)

	if cur == nil || cur.Next == nil {
		return cur
	}

	head = cur.Next

	for cur != nil && cur.Next != nil {
		curNext := cur.Next

		cur.Next = curNext.Next
		curNext.Next = cur
		if prev != nil {
			prev.Next = curNext
		}

		prev = cur
		cur = cur.Next
	}

	return head
}
