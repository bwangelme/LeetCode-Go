package l100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSameTree(t *testing.T) {
	var t1, t2 *TreeNode

	t1 = NewTreeNode(1)
	t1.Left = NewTreeNode(2)
	t1.Right = NewTreeNode(3)

	t2 = NewTreeNode(1)
	t2.Left = NewTreeNode(2)
	t2.Right = NewTreeNode(3)
	assert.True(t, isSameTree(t1, t2))

	t1 = NewTreeNode(1)
	t1.Left = NewTreeNode(2)

	t2 = NewTreeNode(1)
	t2.Right = NewTreeNode(2)
	assert.False(t, isSameTree(t1, t2))


	t1 = NewTreeNode(1)
	t1.Left = NewTreeNode(2)
	t1.Right = NewTreeNode(1)

	t2 = NewTreeNode(1)
	t2.Left = NewTreeNode(1)
	t2.Right = NewTreeNode(2)
	assert.False(t, isSameTree(t1, t2))
}
