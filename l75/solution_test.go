package l75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortColors(t *testing.T) {
	for _, tt := range []struct {
		Nums   []int
		Wanted []int
	}{
		{
			[]int{2, 0, 2, 1, 1, 0},
			[]int{0, 0, 1, 1, 2, 2},
		},
		{
			[]int{2, 0, 1},
			[]int{0, 1, 2},
		},
		{
			[]int{2},
			[]int{2},
		},
	} {
		sortColors(tt.Nums)
		assert.Equal(t, tt.Wanted, tt.Nums)
	}
}
