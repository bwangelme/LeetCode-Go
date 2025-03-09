package l173

import (
	"github.com/bwangelme/LeetCode-Go/lt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BSTIterator(t *testing.T) {
	root := lt.BFSArray2Tree([]interface{}{7, 3, 15, nil, nil, 9, 20})
	i := Constructor(root)
	assert.Equal(t, i.Next(), 3)
	assert.Equal(t, i.Next(), 7)
	assert.Equal(t, i.HasNext(), true)
	assert.Equal(t, i.Next(), 9)
	assert.Equal(t, i.HasNext(), true)
	assert.Equal(t, i.Next(), 15)
	assert.Equal(t, i.HasNext(), true)
	assert.Equal(t, i.Next(), 20)
	assert.Equal(t, i.HasNext(), false)
}
