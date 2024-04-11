package l145

import (
	"github.com/bwangelme/LeetCode-Go/lt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	root := &lt.TreeNode{
		Val: 1,
		Right: &lt.TreeNode{
			Val: 2,
			Left: &lt.TreeNode{
				Val: 3,
			},
		},
	}

	res := postorderTraversal(root)
	assert.Equal(t, []int{3, 2, 1}, res)
}
