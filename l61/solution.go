package l61

import "github.com/bwangelme/LeetCode-Go/lt"

/*
## 题目描述

给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]

## 解题思路

 1. 将链表拼成一个环
 2. 找到链表的尾节点，链表旋转 k ，就是将 n-k 节点，移动 k 次，移动到 n 的位置上。
    2.1 如果 k > n ，那么就是移动 k%n 次
    2.2 我们要找到 n-k 节点，就需要将尾节点 tail 向前移动 n-k 次
 3. 将找到的 n-k 节点断开，n-k+1 就是新的头节点

## 复杂度分析

N = 链表长度

空间复杂度 O(1)

时间复杂度

1. 遍历一边链表
2. 往后遍历 (n - k%n) 个节点

因为 (n - k%n) < n, 因此总的时间复杂度是 O(N)
*/
func rotateRight(head *lt.ListNode, k int) *lt.ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	n := 1
	iter := head
	for iter.Next != nil {
		iter = iter.Next
		n++
	}

	iter.Next = head
	add := n - (k % n)
	for add > 0 {
		iter = iter.Next
		add--
	}

	res := iter.Next
	iter.Next = nil
	return res

}
