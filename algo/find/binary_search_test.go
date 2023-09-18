package find

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BinarySearchFist(t *testing.T) {
	for _, tt := range []struct {
		Arr    []int64
		Num    int64
		Wanted int
	}{
		{
			[]int64{1, 7, 7, 7, 7, 8},
			7,
			1,
		},
		{
			[]int64{7, 7, 7, 7, 7, 7},
			7,
			0,
		},
		{
			[]int64{7, 7, 7, 7, 7, 7},
			8,
			-1,
		},
	} {
		assert.Equal(t, tt.Wanted, BinarySearchFirst(tt.Arr, tt.Num))
		assert.Equal(t, tt.Wanted, BinarySearchFirstV2(tt.Arr, tt.Num))
	}
}

func Test_BinarySearchLast(t *testing.T) {
	for _, tt := range []struct {
		Arr    []int64
		Num    int64
		Wanted int
	}{
		{
			[]int64{1, 7, 7, 7, 7, 8},
			7,
			4,
		},
		{
			[]int64{7, 7, 7, 7, 7, 7},
			7,
			5,
		},
		{
			[]int64{7, 7, 7, 7, 7, 7},
			8,
			-1,
		},
	} {
		assert.Equal(t, tt.Wanted, BinarySearchLast(tt.Arr, tt.Num))
		assert.Equal(t, tt.Wanted, BinarySearchLastV2(tt.Arr, tt.Num))
	}
}
