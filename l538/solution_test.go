package l538

import (
	"github.com/bwangelme/LeetCode-Go/lt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_convertBST(t *testing.T) {
	root := lt.BFSArray2Tree([]interface{}{4, 1, 6, 0, 2, 5, 7, nil, nil, nil, 3, nil, nil, nil, 8})
	root = ConvertBST(root)
	lt.PrintTree(root, os.Stdout)
	assert.Equal(t, root.Val, 30)
}
