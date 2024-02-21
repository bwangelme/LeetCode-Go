package l739

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	for _, tt := range []struct {
		Input  []int
		Wanted []int
	}{
		{
			[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			[]int{30, 40, 50, 60}, []int{1, 1, 1, 0},
		},
		{
			[]int{30, 60, 90}, []int{1, 1, 0},
		},
	} {
		assert.Equal(t, tt.Wanted, dailyTemperatures(tt.Input))
	}
}
