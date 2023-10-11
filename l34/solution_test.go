package l34

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchRange(t *testing.T) {
	assert.Equal(t, searchRange([]int{5, 7, 7, 8, 8, 10}, 8), []int{3, 4})
	assert.Equal(t, searchRange([]int{5, 7, 7, 8, 8, 10}, 6), []int{-1, -1})
	assert.Equal(t, searchRange([]int{}, 0), []int{-1, -1})
}
