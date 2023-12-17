package l19

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
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

func PrintList(head *ListNode) string {
	var b = strings.Builder{}
	var dummy = &ListNode{
		Next: head,
	}
	var current = dummy

	for current.Next != nil {
		current = current.Next
		b.WriteString(fmt.Sprint(current.Val))
		if current.Next != nil {
			b.WriteString(" -> ")
		}
	}

	return b.String()
}

func Test_removeNthFromEnd(t *testing.T) {
	for _, tt := range []struct {
		arr    []int
		n      int
		wanted string
	}{
		{
			[]int{1},
			1,
			"",
		},
		{
			[]int{1, 2, 3, 4, 5},
			2,
			"1 -> 2 -> 3 -> 5",
		},
	} {
		head := NewList(tt.arr)
		head = removeNthFromEnd(head, tt.n)
		assert.Equal(t, tt.wanted, PrintList(head))
	}

}
