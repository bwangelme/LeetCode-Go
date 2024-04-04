package l513

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findBottomLeftValue(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	res := findBottomLeftValue(root)
	assert.Equal(t, res, 5)
}
