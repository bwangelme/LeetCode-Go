package l814

import (
	"github.com/bwangelme/LeetCode-Go/lt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pruneTree(t *testing.T) {
	root := &lt.TreeNode{
		Val: 1,
		Right: &lt.TreeNode{
			Val: 0,
			Left: &lt.TreeNode{
				Val: 0,
			},
			Right: &lt.TreeNode{
				Val: 1,
			},
		},
	}

	root = pruneTree(root)
	assert.NotNil(t, root)
	assert.Equal(t, root.Right.Right.Val, 1)
}
