package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MergeSort(t *testing.T) {
	for _, tt := range []struct {
		Val    []int64
		Wanted []int64
	}{
		{
			[]int64{11, 8, 3, 9, 7, 1, 2, 5},
			[]int64{1, 2, 3, 5, 7, 8, 9, 11},
		},
		{
			[]int64{42},
			[]int64{42},
		},
	} {
		assert.Equal(t, tt.Wanted, MergeSort(tt.Val))
	}
}
