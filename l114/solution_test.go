package l114

import (
	"testing"

	"github.com/bwangelme/LeetCode-Go/lt"
	"github.com/stretchr/testify/assert"
)

func Test_flatten(t *testing.T) {
	root := &lt.TreeNode{
		Val: 1,
		Left: &lt.TreeNode{
			Val: 2,
			Left: &lt.TreeNode{
				Val: 3,
			},
			Right: &lt.TreeNode{
				Val: 4,
			},
		},
		Right: &lt.TreeNode{
			Val: 5,
			Right: &lt.TreeNode{
				Val: 6,
			},
		},
	}
	flatten(root)
	//lt.PrintTree(root, os.Stdout)
	assert.Equal(t, root.Right.Right.Right.Right.Val, 5)
}

func Test_flatten_v2(t *testing.T) {
	root := &lt.TreeNode{
		Val: 1,
		Left: &lt.TreeNode{
			Val: 2,
			Left: &lt.TreeNode{
				Val: 3,
			},
		},
	}
	flatten(root)
	//lt.PrintTree(root, os.Stdout)
	assert.Equal(t, root.Right.Right.Val, 3)
}
