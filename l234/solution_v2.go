package l234

/*
## 题目

给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

提示：

- 链表中节点数目在范围[1, 10^5] 内
- 0 <= Node.val <= 9

## 解题思路

1. 将链表分成前后两部分，如果是奇数个节点，忽略中间节点
2. 将前半部分反转
3. 比较前后两部分是否相同

## 复杂度分析

N 表示链表的长度
时间复杂度: O(N)
空间复杂度: O(1)
*/
func isPalindromeV2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	var (
		slow = head
		fast = head.Next
	)

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// slow 指向前半部分尾部，secondHalf 指向后半部分头部
	// 如果是奇数个节点，忽略中间的数
	secondHalf := slow.Next
	if fast.Next != nil {
		secondHalf = slow.Next.Next
	}
	slow.Next = nil

	return equal(reverseList(head), secondHalf)
}

func equal(node1, node2 *ListNode) bool {
	for node1 != nil && node2 != nil {
		if node1.Val != node2.Val {
			return false
		}

		node1 = node1.Next
		node2 = node2.Next
	}

	return node1 == nil && node2 == nil
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
