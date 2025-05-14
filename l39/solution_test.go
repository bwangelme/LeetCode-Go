package l39

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum(t *testing.T) {
	nums := []int{2, 3, 6, 7}
	target := 7
	result := combinationSum(nums, target)
	assert.Equal(t, result, [][]int{
		{7},
		{2, 2, 3},
	})
}
