package lcr053

import (
	"github.com/bwangelme/LeetCode-Go/lt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_inorderSuccessor(t *testing.T) {
	root := lt.BFSArray2Tree([]interface{}{2, 1, 3})
	p := &lt.TreeNode{
		Val: 1,
	}
	assert.Equal(t, inorderSuccessor(root, p).Val, 2)
}
