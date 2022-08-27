package l704

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_search(t *testing.T) {
	for _, tt := range []struct {
		nums   []int
		target int

		wanted int
	}{
		{
			[]int{-1, 0, 3, 5, 9, 12},
			9,
			4,
		},
		{
			[]int{-1, 0, 3, 5, 9, 12},
			2,
			-1,
		},
		{
			[]int{5},
			5,
			0,
		},
		{
			[]int{5},
			3,
			-1,
		},
	} {
		assert.Equal(t, tt.wanted, search(tt.nums, tt.target))
	}

}
