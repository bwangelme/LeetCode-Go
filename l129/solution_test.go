package l129

import (
	"github.com/bwangelme/LeetCode-Go/lt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sumNumbers(t *testing.T) {
	root := &lt.TreeNode{
		Val: 1,
		Left: &lt.TreeNode{
			Val: 2,
		},
		Right: &lt.TreeNode{
			Val: 3,
		},
	}
	assert.Equal(t, 25, sumNumbers(root))
}
