package l24

import (
	"testing"

	"github.com/bwangelme/LeetCode-Go/lt"
	"github.com/stretchr/testify/assert"
)

func Test_swapPairs(t *testing.T) {
	head := lt.CreateList([]int{1, 2, 3, 4})
	head = swapPairs(head)
	assert.Equal(t, "2 1 4 3", head.String())

	head = lt.CreateList([]int{1, 2, 3, 4, 5})
	head = swapPairs(head)
	assert.Equal(t, "2 1 4 3 5", head.String())
}
