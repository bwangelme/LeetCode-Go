package l4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMedianSortedArrays(t *testing.T) {
	for _, tt := range []struct{
		arr1 []int
		arr2 []int
		expected float64
	}{
		{
			arr1: []int{1, 3},
			arr2: []int{2},
			expected: 2.0,
		},
		{
			arr1: []int{1, 2},
			arr2: []int{3, 4},
			expected: 2.5,
		},
		{
			arr1: []int{0, 0},
			arr2: []int{0, 0},
			expected: 0,
		},
		{
			arr1: []int{1, 3},
			arr2: []int{2},
			expected: 2,
		},
		{
			arr1: []int{},
			arr2: []int{1},
			expected: 1,
		},
		{
			arr1: []int{2},
			arr2: []int{},
			expected: 2,
		},
	}{
		assert.Equal(t, tt.expected, findMedianSortedArrays(tt.arr1, tt.arr2))
	}
}


func Test_findKth(t *testing.T) {
	arr1 := []int{1, 3, 9}
	arr2 := []int{0, 2, 4, 7}
	assert.Equal(t, 4, findKth(arr1, len(arr1), arr2, len(arr2), 5))
}
