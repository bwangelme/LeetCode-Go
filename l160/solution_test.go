package l160

import (
	"github.com/bwangelme/LeetCode-Go/lt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	var (
		pA = lt.CreateList([]int{1, 2, 3})
		pB = lt.CreateList([]int{4, 5, 6})
	)
	assert.Nil(t, getIntersectionNode(pA, pB))

	pA = lt.CreateList([]int{1, 2, 3, 4, 5, 6})
	pB = lt.CreateList([]int{1, 2})
	pB.Next.Next = pA.Next.Next
	assert.Equal(t, getIntersectionNode(pA, pB).Val, 3)
}
