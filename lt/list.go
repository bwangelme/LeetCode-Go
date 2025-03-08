package lt

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

	// 去除末尾的 ->
	res = res[:len(res)-1]

	return res
}

func CreateList(nums []int) *ListNode {
	head := NewListNode(nums[0], nil)
	prevNode := head

	for _, val := range nums[1:] {
		node := NewListNode(val, nil)
		prevNode.Next = node
		prevNode = node
	}

	return head
}

func NewListNode(Val int, Next *ListNode) *ListNode {
	return &ListNode{
		Val:  Val,
		Next: Next,
	}
}
