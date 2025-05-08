package l148

import (
	"github.com/bwangelme/LeetCode-Go/lt"
)

/*

## 题目描述

给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

## 解题思路

使用归并排序的方式进行排序

递归地将链表二分，然后将二分的链表有序合并起来。合并的时候，不需要创建额外的数组存储结果，没有像归并排序数组一样申请额外的空间。

## 复杂度分析

N = 链表长度

时间复杂度: O(NlogN), 将链表不断二分进行排序, 一共进行了 logN 次调用，每次合并的时候，会遍历数组，因此是 NlogN
空间复杂度: 进行了 O(logN) 次的函数调用，因此空间复杂度是 O(logN)
*/

func sortList(head *lt.ListNode) *lt.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	h1 := head
	h2 := split(head)

	h1 = sortList(h1)
	h2 = sortList(h2)

	return merge(h1, h2)
}

func split(head *lt.ListNode) *lt.ListNode {
	var (
		slow = head
		fast = head.Next
	)

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := slow.Next
	slow.Next = nil
	return second
}

func merge(head1, head2 *lt.ListNode) *lt.ListNode {
	var (
		dummy = &lt.ListNode{}
		cur   = dummy
	)

	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			cur.Next = head1
			head1 = head1.Next
		} else {
			cur.Next = head2
			head2 = head2.Next
		}

		cur = cur.Next
	}

	if head1 == nil {
		cur.Next = head2
	} else {
		cur.Next = head1
	}

	return dummy.Next
}
