package l515

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_largestValues(t *testing.T) {
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

	res := largestValues(root)
	assert.Equal(t, res, []int{1, 3, 9})
}
