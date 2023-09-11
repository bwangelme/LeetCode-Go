package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TopK(t *testing.T) {
	for _, tt := range []struct {
		Arr    []int64
		K      int
		Wanted int64
	}{
		{
			[]int64{6, 1, 3, 5, 7, 2, 4, 9, 11, 8},
			3,
			3,
		},
		{
			[]int64{8},
			2,
			-1,
		},
	} {
		assert.Equal(t, tt.Wanted, TopK(tt.Arr, tt.K))
	}
}
