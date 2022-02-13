package l102

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newNode(val int) *TreeNode {
	return &TreeNode{
		Left:  nil,
		Right: nil,
		Val:   val,
	}
}

func Test_levelOrder_1(t *testing.T) {
	root := newNode(3)
	root.Left = newNode(9)
	root.Right = newNode(20)
	root.Right.Left = newNode(15)
	root.Right.Right = newNode(7)

	res := levelOrder(root)
	assert.Equal(t, res, [][]int{{3}, {9, 20}, {15, 7}})
}

func Test_levelOrder_2(t *testing.T) {
	root := newNode(1)

	res := levelOrder(root)

	assert.Equal(t, res, [][]int{{1}})
}

func Test_levelOrder_3(t *testing.T) {
	res := levelOrder(nil)

	assert.Equal(t, res, [][]int{})
}
