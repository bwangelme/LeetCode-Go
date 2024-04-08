package l144

import (
	"github.com/bwangelme/LeetCode-Go/lt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	root := &lt.TreeNode{
		Val: 1,
		Right: &lt.TreeNode{
			Val: 2,
			Left: &lt.TreeNode{
				Val: 3,
			},
		},
	}

	res := preorderTraversal(root)
	assert.Equal(t, []int{1, 2, 3}, res)
}
