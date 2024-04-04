package l199

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rightSideView(t *testing.T) {
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

	res := rightSideView(root)
	assert.Equal(t, []int{1, 2, 9}, res)
}
