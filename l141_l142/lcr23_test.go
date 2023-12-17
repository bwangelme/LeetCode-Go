package l141_l142

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	//A: 4 -> 1 -> 8 -> 4 -> 5
	//            /|\
	//             |
	//B: 5 -> 0 -> 1
	headA := NewList([]int{4, 1, 8, 4, 5}, -1)
	headB := NewList([]int{5, 0, 1}, -1)
	headB.Next.Next.Next = headA.Next.Next
	assert.Equal(t, 8, getIntersectionNode(headA, headB).Val)

	//A: 4 -> 1 -> 8 -> 4 -> 5
	//B: 5 -> 0 -> 1
	headA = NewList([]int{4, 1, 8, 4, 5}, -1)
	headB = NewList([]int{5, 0, 1}, -1)
	assert.Nil(t, getIntersectionNode(headA, headB))
}
