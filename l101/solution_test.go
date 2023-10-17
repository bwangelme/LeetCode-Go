package l101

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isSymmetric(t *testing.T) {
	var root *TreeNode

	//root = NewTreeNode("1")
	//root.Left = NewTreeNode("2")
	//root.Right = NewTreeNode("2")
	//root.Left.Left = NewTreeNode("3")
	//root.Left.Right = NewTreeNode("4")
	//root.Right.Left = NewTreeNode("4")
	//root.Right.Right = NewTreeNode("3")
	//assert.Equal(t, true, isSymmetric(root))

	root = NewTreeNode("1")
	root.Left = NewTreeNode("2")
	root.Right = NewTreeNode("2")
	root.Left.Left = NewTreeNode("2")
	root.Left.Right = nil
	root.Right.Left = NewTreeNode("2")
	root.Right.Right = nil
	assert.Equal(t, false, isSymmetric(root))
}
