package l876

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_middleNode(t *testing.T) {
	for _, tt := range []struct {
		vals   []int
		wanted int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 4},
		{[]int{1, 2, 3, 4, 5}, 3},
	} {
		head := NewList(tt.vals)
		res := middleNode(head)

		assert.Equal(t, tt.wanted, res.Val)
	}
}
