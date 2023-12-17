package lcr140

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewList(arr []int) *ListNode {
	var (
		dummy   = &ListNode{}
		current = dummy
	)

	for i := 0; i < len(arr); i++ {
		node := &ListNode{
			Val: arr[i],
		}
		current.Next = node
		current = current.Next
	}

	return dummy.Next
}

func Test_trainingPlan(t *testing.T) {
	for _, tt := range []struct {
		Arr    []int
		Cnt    int
		Wanted int
	}{
		{
			[]int{2, 4, 7, 8}, 1, 8,
		},
	} {
		head := NewList(tt.Arr)
		res := trainingPlan(head, tt.Cnt)
		assert.Equal(t, tt.Wanted, res.Val)
	}
}
