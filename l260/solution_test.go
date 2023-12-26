package l260

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	for _, tt := range []struct {
		Input  []int
		Wanted []int
	}{
		{
			[]int{1, 2, 1, 3, 2, 5},
			[]int{3, 5},
		},
		{
			[]int{0, 1},
			[]int{0, 1},
		},
	} {
		res := singleNumber(tt.Input)
		sort.Ints(res)
		assert.Equal(t, tt.Wanted, res)
	}
}
