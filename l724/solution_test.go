package l724

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pivotIndex(t *testing.T) {
	for _, tt := range []struct {
		Arr    []int
		Wanted int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{1, 2, 3}, -1},
	} {
		assert.Equal(t, tt.Wanted, pivotIndex(tt.Arr))
	}
}
