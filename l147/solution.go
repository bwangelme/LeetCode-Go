package l147

import (
	"github.com/bwangelme/LeetCode-Go/lt"
)

/*
## 题目描述

对链表进行插入排序

## 解题思路

lastSorted 指向排序区域的最后一个数字，每次 cur 通过 lastSorted->Next 来获取

## 复杂度分析

N = 链表长度
- 空间复杂度

没有申请新的空间，O(1)

- 时间复杂度
cur 等于遍历链表中的每个节点，进行插入
每次找 cur 的 prev ，都要重头开始找

因此整个遍历次数是 1+2+...N, 因此时间复杂度是 O(N^2)
*/
func insertionSortList(head *lt.ListNode) *lt.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var (
		dummy = &lt.ListNode{
			Next: head,
		}
		cur        = head.Next
		lastSorted = head
	)

	for cur != nil {
		if lastSorted.Val <= cur.Val {
			lastSorted = lastSorted.Next
		} else {
			// 寻找 cur 的插入位置，prev, cur 放在 prev 之后
			// prev 是从头开始找的, 然后找到第一个比 cur 大的节点
			prev := dummy
			for prev.Next.Val <= cur.Val {
				prev = prev.Next
			}
			lastSorted.Next = cur.Next
			cur.Next = prev.Next
			prev.Next = cur
		}
		cur = lastSorted.Next
	}

	return dummy.Next
}
