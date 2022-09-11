package l876

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l ListNode) String() string {
	return fmt.Sprintf("%v", l.Val)
}

func NewNode(val int, next *ListNode) *ListNode {
	return &ListNode{
		Val:  val,
		Next: next,
	}
}

func NewList(vals []int) *ListNode {
	var nodeMap = make(map[int]*ListNode)

	for idx, val := range vals {
		node := NewNode(val, nil)
		nodeMap[idx] = node

		if idx > 0 {
			nodeMap[idx-1].Next = node
		}
	}

	return nodeMap[0]
}

/*
## 解题思路

题目说需要返回链表的中间节点，使用快慢指针同时遍历，快指针到头时，慢指针就在中间

## 复杂度分析

假设链表的长度为 N

时间复杂度: O(N/2) => O(N)
空间复杂度: O(1)，共申请了两个指针的空间
*/
func middleNode(head *ListNode) *ListNode {
	// 题目中说了非空链表，所以不判断 head == nil 了
	var (
		slow = head
		fast = head
	)

	for {
		// 链表长度是奇数个
		if fast.Next == nil {
			return slow
		}

		slow = slow.Next
		fast = fast.Next.Next

		// 链表长度是偶数个
		if fast == nil {
			return slow
		}
	}
}
