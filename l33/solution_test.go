package l33

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMin(t *testing.T) {
	assert.Equal(t, 4, findMinIdx([]int{4, 5, 6, 7, 0, 1, 2}))
	assert.Equal(t, 6, findMinIdx([]int{1, 2, 4, 5, 6, 7, 0}))
}

func Test_binSearch(t *testing.T) {
	assert.Equal(t, -1, binSearch([]int{4, 5, 6, 7}, 0, 3, 8))
	assert.Equal(t, 2, binSearch([]int{4, 5, 6, 7}, 0, 3, 6))
}

func Test_search(t *testing.T) {
	assert.Equal(t, 4, search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	assert.Equal(t, -1, search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
}
