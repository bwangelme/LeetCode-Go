package l19

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		dummy = &ListNode{
			Next: head,
		}
		first  = head
		second = dummy
	)

	for ; n > 0; n-- {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next

	return dummy.Next
}
