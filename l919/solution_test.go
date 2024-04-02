package l919

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CBTInserter(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	obj := Constructor(root)
	v := obj.Insert(3)
	assert.Equal(t, 1, v)
	v = obj.Insert(4)
	assert.Equal(t, 2, v)
	root = obj.Get_root()
	assert.Equal(t, 3, root.Right.Val)

}
