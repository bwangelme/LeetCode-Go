package l15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_threeSum(t *testing.T) {
	for _, tt := range []struct {
		Input  []int
		Wanted [][]int
	}{
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
	} {
		assert.Equal(t, tt.Wanted, threeSum(tt.Input))
	}
}
