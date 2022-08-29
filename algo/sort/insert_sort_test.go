package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_InsertSort(t *testing.T) {
	input := []int{3, 44, 38, 5, 26, 27}
	wanted := []int{3, 5, 26, 27, 38, 44}
	res := InsertSort(input)
	assert.Equal(t, wanted, res)

	i := []int32{2, 87, 39, 49, 34, 62, 53, 6, 44, 98}
	w := []int32{2, 6, 34, 39, 44, 49, 53, 62, 87, 98}
	r := InsertSort(i)
	assert.Equal(t, w, r)
}

func Test_InsertSortWithSentinel(t *testing.T) {
	input := []int{3, 44, 38, 5, 26, 27}
	wanted := []int{3, 5, 26, 27, 38, 44}
	res := InsertSortWithSentinel(input)
	assert.Equal(t, wanted, res)

	i := []int32{2, 87, 39, 49, 34, 62, 53, 6, 44, 98}
	w := []int32{2, 6, 34, 39, 44, 49, 53, 62, 87, 98}
	r := InsertSortWithSentinel(i)
	assert.Equal(t, w, r)
}
