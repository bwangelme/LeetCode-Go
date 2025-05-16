package l147

import (
	"testing"

	"github.com/bwangelme/LeetCode-Go/lt"
	"github.com/stretchr/testify/assert"
)

func Test_insertionSortList(t *testing.T) {
	var head *lt.ListNode = nil

	head = lt.CreateList([]int{4, 2, 1, 3})
	head = insertionSortList(head)
	assert.Equal(t, head.String(), "1 2 3 4")

	head = lt.CreateList([]int{6, 5, 3, 1, 8, 7, 2, 4})
	head = insertionSortList(head)
	assert.Equal(t, "1 2 3 4 5 6 7 8", head.String())
}
