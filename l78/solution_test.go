package l78

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subsets(t *testing.T) {
	nums := []int{1, 2, 3}
	res := subsets(nums)
	assert.Equal(t, res, [][]int{
		{},
		{3}, {2},
		{2, 3}, {1}, {1, 3}, {1, 2},
		{1, 2, 3},
	})
}
