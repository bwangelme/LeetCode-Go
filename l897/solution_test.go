package l897

import (
	"github.com/bwangelme/LeetCode-Go/lt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_increasingBST(t *testing.T) {
	root := lt.BFSArray2Tree([]interface{}{5, 3, 6, 2, 4, nil, 8, 1, nil, nil, nil, 7, 9})
	node := increasingBST(root)
	lt.PrintTree(node, os.Stdout)
	assert.Equal(t, node.Val, 1)
}

func Test_increasingBSTV2(t *testing.T) {
	root := lt.BFSArray2Tree([]interface{}{5, 3, 6, 2, 4, nil, 8, 1, nil, nil, nil, 7, 9})
	node := increasingBSTV2(root)
	lt.PrintTree(node, os.Stdout)
	assert.Equal(t, node.Val, 1)
}
