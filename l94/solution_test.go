package l94

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	var tr *TreeNode

	tr = NewTreeNode(1)
	tr.Right = NewTreeNode(2)
	tr.Right.Left = NewTreeNode(3)

	//assert.Equal(t, []int{1, 3, 2}, inorderTraversal_recursion(tr))
	assert.Equal(t, []int{1, 3, 2}, inorderTraversal(tr))
}
