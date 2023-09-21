package find

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BinarySearchFirstGt(t *testing.T) {
	for _, tt := range []struct {
		Arr    []int64
		Num    int64
		Wanted int
	}{
		{
			[]int64{-12, 2, 3, 4, 9, 132},
			9,
			5,
		},
		{
			[]int64{-12, 2, 3, 4, 9, 132},
			-13,
			0,
		},
		{
			[]int64{-12, 2, 3, 4, 9, 132},
			133,
			-1,
		},
	} {
		assert.Equal(t, tt.Wanted, BinarySearchFirstGt(tt.Arr, tt.Num))
	}
}

func Test_BinarySearchFirstLt(t *testing.T) {
	for _, tt := range []struct {
		Arr    []int64
		Num    int64
		Wanted int
	}{
		{
			[]int64{-12, 2, 3, 4, 9, 132},
			133,
			5,
		},
		{
			[]int64{-12, 2, 3, 4, 9, 132},
			5,
			3,
		},
		{
			[]int64{-12, 2, 3, 4, 9, 132},
			-13,
			-1,
		},
	} {
		assert.Equal(t, tt.Wanted, BinarySearchLastLt(tt.Arr, tt.Num))
	}
}
