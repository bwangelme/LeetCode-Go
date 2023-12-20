package l445

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseList(t *testing.T) {
	l1 := createList([]int{3, 2, 1})
	l2 := reverseList(l1)
	assert.Equal(t, "1 2 3", l2.String())
}

func Test_addTwoNumbers(t *testing.T) {
	for _, tt := range []struct {
		a1     []int
		a2     []int
		Wanted string
	}{
		{
			a1:     []int{7, 2, 4, 3},
			a2:     []int{5, 6, 4},
			Wanted: "7 8 0 7",
		},
		{
			a1:     []int{2, 4, 3},
			a2:     []int{5, 6, 4},
			Wanted: "8 0 7",
		},
		{
			a1:     []int{0},
			a2:     []int{0},
			Wanted: "0",
		},
		{
			a1:     []int{5},
			a2:     []int{5},
			Wanted: "1 0",
		},
	} {
		l1 := createList(tt.a1)
		l2 := createList(tt.a2)
		res := addTwoNumbers(l1, l2)
		assert.Equal(t, tt.Wanted, res.String())
	}
}
