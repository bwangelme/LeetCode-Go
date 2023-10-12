package l81

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_search(t *testing.T) {
	for _, tt := range []struct {
		Nums   []int
		Target int
		Wanted bool
	}{
		{
			[]int{1, 0, 1, 1, 1},
			0,
			true,
		},
		{
			[]int{2, 5, 6, 0, 0, 1, 2},
			3,
			false,
		},
		{
			[]int{2, 5, 6, 0, 0, 1, 2},
			0,
			true,
		},
	} {
		assert.Equal(t, tt.Wanted, search(tt.Nums, tt.Target))
	}
}
