package l437

import (
	"github.com/bwangelme/LeetCode-Go/lt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pathSum(t *testing.T) {
	var input = []interface{}{10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1}
	var targetSum = 8
	var wanted = 3

	root := lt.BFSArray2Tree(input)
	assert.Equal(t, pathSum(root, targetSum), wanted)
}
