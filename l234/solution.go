package l234

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  rune
	Next *ListNode
}

func (l ListNode) String() string {
	return fmt.Sprintf("%v", string(l.Val))
}

func NewListNode(r rune, next *ListNode) *ListNode {
	return &ListNode{
		Val:  r,
		Next: next,
	}
}

func PrintList(head *ListNode) string {
	var b = strings.Builder{}
	var current *ListNode = head

	for {
		b.WriteRune(current.Val)
		if current.Next != nil {
			b.WriteString(" -> ")
			current = current.Next
		} else {
			break
		}
	}

	return b.String()
}

func NewList(str string) *ListNode {
	var (
		head    *ListNode
		current = head
	)

	for _, r := range str {
		if head == nil {
			head = NewListNode(r, nil)
			current = head
		} else {
			current.Next = NewListNode(r, nil)
			current = current.Next
		}
	}

	return head
}

/**
## 题目

给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false

## 解题思路

1. 定义快慢两个指针指向 head，慢指针每次前进一步，快指针每次前进两步
2. 慢指针每次将 next 改成 pre
	next = slow.next
	slow.next
	prev
	slow = next
3. 用 prev 记录上一次慢指针遍历的节点
4. fast == null 的时候，slow 在中间，如果是长度奇数，需要额外+1，忽略最中间的节点
5.将 slow 和 prev 进行比较

## 注意事项

1. len == 0 || len == 1 的链表都是回文
2. abcba 这种奇数长度的 list，忽略最中间的节点

## 时间复杂度

O(N)

*/

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	var (
		fast           = head
		slow           = head
		prev *ListNode = nil
	)

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next

		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	if fast != nil {
		slow = slow.Next
	}

	for slow != nil && prev != nil {
		if slow.Val != prev.Val {
			return false
		}
		slow = slow.Next
		prev = prev.Next
	}

	return true
}
