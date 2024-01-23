package l304

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NumMatrix(t *testing.T) {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	obj := Constructor(matrix)

	assert.Equal(t, 8, obj.SumRegion(2, 1, 4, 3))
	assert.Equal(t, 11, obj.SumRegion(1, 1, 2, 2))
	assert.Equal(t, 12, obj.SumRegion(1, 2, 2, 4))
}
