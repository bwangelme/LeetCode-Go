package l23

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func makeList(arr []int) *ListNode {
	var (
		head    = &ListNode{}
		current = head
	)
	for _, val := range arr {
		current.Next = &ListNode{
			Val:  val,
			Next: nil,
		}
		current = current.Next
	}

	return head.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}

	fmt.Println()
}

func listEqual(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 == nil && l2 == nil {
		return true
	}
	return false
}

func Test_mergeKLists(t *testing.T) {
	var lists = []*ListNode{
		makeList([]int{1, 4, 5}),
		makeList([]int{1, 3, 4}),
		makeList([]int{2, 6}),
	}

	res := mergeKLists(lists)
	assert.True(t, listEqual(makeList([]int{1, 1, 2, 3, 4, 4, 5, 6}), res))

	lists = []*ListNode{
		makeList([]int{2}),
		nil,
		makeList([]int{-1}),
	}

	res = mergeKLists(lists)
	assert.True(t, listEqual(makeList([]int{-1, 2}), res))
}
