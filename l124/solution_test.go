package l124

import (
	"github.com/bwangelme/LeetCode-Go/lt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxPathSum(t *testing.T) {
	var input = []interface{}{2, -1}
	var wanted = 2

	root := lt.BFSArray2Tree(input)
	assert.Equal(t, maxPathSum(root), wanted)
}
