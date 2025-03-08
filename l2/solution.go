package l2

import (
	"github.com/bwangelme/LeetCode-Go/lt"
)

// 设 l1 的长度为 M，l2 的长度为N，(N >= M)。
// 时间复杂度为 O(N)
// 空间复杂度为 O(N)，不考虑 l1 和 l2 已经占用的空间，若考虑的话，则是 O(N + M)
func addTwoNumbers(l1 *lt.ListNode, l2 *lt.ListNode) *lt.ListNode {
	var (
		head, prevNode *lt.ListNode
		carry, num     int
	)

	for {
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
		num = 0

		if l1 != nil {
			num += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			num += l2.Val
			l2 = l2.Next
		}

		num += carry
		carry = 0

		if num >= 10 {
			num -= 10
			carry = 1
		}

		node := lt.NewListNode(num, nil)

		if head == nil {
			head = node
		}

		if prevNode != nil {
			prevNode.Next = node
		}

		prevNode = node
	}

	return head
}
