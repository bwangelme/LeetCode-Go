package l18

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fourSum(t *testing.T) {
	for _, tt := range []struct {
		nums    []int
		target  int
		answers [][]int
	}{
		{
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
			answers: [][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			nums:   []int{-1, 0, 1, 2, -1, -4},
			target: -1,
			answers: [][]int{
				{-4, 0, 1, 2},
				{-1, -1, 0, 1},
			},
		},
	} {
		result := fourSum(tt.nums, tt.target)
		assert.Equal(t, len(result), len(tt.answers))
		for idx, res := range tt.answers {
			assert.InDeltaSlice(t, res, result[idx], 0)
		}
	}
}
