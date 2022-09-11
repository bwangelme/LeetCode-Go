package l21

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeTwoLists(t *testing.T) {
	for _, tt := range []struct {
		l1     []int
		l2     []int
		wanted []int
	}{
		{
			[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4},
		},
		{
			[]int{}, []int{}, []int{},
		},
		{
			[]int{}, []int{0}, []int{0},
		},
	} {
		l1 := NewList(tt.l1)
		l2 := NewList(tt.l2)
		res := mergeTwoLists(l1, l2)

		assert.Equal(t, tt.wanted, ListToArray(res))
	}
}
