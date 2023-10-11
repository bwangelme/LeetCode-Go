package l74

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchFirstGtRow(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	row, _ := searchFirstGteRow(matrix, 3)
	assert.Equal(t, 0, row)
}

func Test_searchMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	assert.Equal(t, true, searchMatrix(matrix, 3))
	assert.Equal(t, false, searchMatrix(matrix, 13))

	matrix = [][]int{
		{1},
	}
	assert.Equal(t, true, searchMatrix(matrix, 1))
}
