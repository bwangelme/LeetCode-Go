package l110

import (
	"github.com/bwangelme/LeetCode-Go/lt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_isBalanced(t *testing.T) {
	root := lt.BFSArray2Tree([]interface{}{1, 2, 2, 3, nil, nil, 3, 4, nil, nil, 4})
	lt.PrintTree(root, os.Stdout)
	assert.Equal(t, false, isBalanced(root))
}
