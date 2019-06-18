package l2

import "fmt"

/**
 * Definition for singly-linked list.
 */
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

func newListNode(Val int, Next *ListNode) *ListNode {
	return &ListNode{
		Val:  Val,
		Next: Next,
	}
}

func listToNum(numList *ListNode) (result uint64) {
	var exponential = 1

	for numList != nil {
		result = uint64(numList.Val*exponential) + result
		numList = numList.Next
		exponential *= 10
	}

	return result
}

func numToList(num uint64) (list *ListNode) {
	list = newListNode(int(num%10), nil)
	num = num / 10
	preNode := list

	for num != 0 {
		node := newListNode(int(num%10), nil)
		preNode.Next = node
		preNode = node
		num = num / 10
	}

	return list
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		n1 = listToNum(l1)
		n2 = listToNum(l2)
	)

	res := n1 + n2

	return numToList(res)
}
