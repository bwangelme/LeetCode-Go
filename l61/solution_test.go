package l61

import (
	"testing"

	"github.com/bwangelme/LeetCode-Go/lt"
	"github.com/stretchr/testify/assert"
)

func Test_rotateRight(t *testing.T) {
	head := lt.CreateList([]int{1, 2, 3, 4, 5})
	head = rotateRight(head, 2)
	assert.Equal(t, "4 5 1 2 3", head.String())

	head = lt.CreateList([]int{0, 1, 2})
	head = rotateRight(head, 4)
	assert.Equal(t, "2 0 1", head.String())

	head = lt.CreateList([]int{1})
	head = rotateRight(head, 0)
	assert.Equal(t, "1", head.String())

	head = lt.CreateList([]int{1, 2})
	head = rotateRight(head, 1)
	assert.Equal(t, "2 1", head.String())
}
