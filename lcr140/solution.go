package lcr140

type ListNode struct {
	Val  int
	Next *ListNode
}

func trainingPlan(head *ListNode, cnt int) *ListNode {
	var (
		first  = head
		second = head
	)

	for i := 1; i < cnt; i++ {
		first = first.Next
	}

	for first.Next != nil {
		first = first.Next
		second = second.Next
	}

	return second
}
