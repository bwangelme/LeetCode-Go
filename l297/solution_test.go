package l297

import (
	"github.com/bwangelme/LeetCode-Go/lt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Codec(t *testing.T) {
	root := &lt.TreeNode{
		Val: 1,
		Left: &lt.TreeNode{
			Val: 2,
		},
		Right: &lt.TreeNode{
			Val: 3,
			Left: &lt.TreeNode{
				Val: 4,
			},
			Right: &lt.TreeNode{
				Val: 5,
			},
		},
	}

	ser := Constructor()
	deser := Constructor()
	data := ser.serialize(root)
	assert.Equal(t, data, "1,2,#,#,3,4,#,#,5,#,#")
	ans := deser.deserialize(data)
	assert.NotNil(t, ans)
	assert.Equal(t, ans.Right.Left.Val, 4)
}
