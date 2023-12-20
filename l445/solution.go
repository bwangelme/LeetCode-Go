package l445

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return fmt.Sprintf("Empty")
	}

	var (
		res  string
		node = l
	)
	for node != nil {
		res += fmt.Sprintf("%d ", node.Val)
		node = node.Next
	}

	res = res[:len(res)-1]

	return res
}

func createList(nums []int) *ListNode {
	head := newListNode(nums[0], nil)
	prevNode := head

	for _, val := range nums[1:] {
		node := newListNode(val, nil)
		prevNode.Next = node
		prevNode = node
	}

	return head
}

func newListNode(Val int, Next *ListNode) *ListNode {
	return &ListNode{
		Val:  Val,
		Next: Next,
	}
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseList(l1)
	l2 = reverseList(l2)

	dummy := newListNode(0, nil)
	cur := dummy
	carry := 0
	for l1 != nil || l2 != nil {
		v1 := 0
		if l1 != nil {
			v1 = l1.Val
		}
		v2 := 0
		if l2 != nil {
			v2 = l2.Val
		}
		val := v1 + v2 + carry
		if val >= 10 {
			val = val - 10
			carry = 1
		} else {
			carry = 0
		}
		cur.Next = newListNode(val, nil)
		cur = cur.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		cur.Next = newListNode(carry, nil)
	}

	res := dummy.Next
	res = reverseList(res)

	return res
}
