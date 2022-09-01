package l141_l142

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasCycle(t *testing.T) {
	for _, tt := range []struct {
		vals   []int
		pos    int
		wanted bool
	}{
		{[]int{3, 2, 0, -4}, 1, true},
		{[]int{1, 2}, 0, true},
		{[]int{1}, -1, false},
	} {
		head := NewList(tt.vals, tt.pos)
		res := hasCycle(head)

		assert.Equal(t, tt.wanted, res)
	}
}

func Test_hasCycleO1(t *testing.T) {
	for _, tt := range []struct {
		vals   []int
		pos    int
		wanted bool
	}{
		{[]int{3, 2, 0, -4}, 1, true},
		{[]int{1, 2}, 0, true},
		{[]int{1}, -1, false},
	} {
		head := NewList(tt.vals, tt.pos)
		res := hasCycleO1(head)

		assert.Equal(t, tt.wanted, res)
	}
}

func Test_detectCycle(t *testing.T) {
	for _, tt := range []struct {
		vals   []int
		pos    int
		wanted int
	}{
		{[]int{3, 2, 0, -4}, 1, 2},
		//{[]int{1, 2}, 0, true},
		//{[]int{1}, -1, false},
	} {
		head := NewList(tt.vals, tt.pos)
		res := detectCycle(head)

		assert.Equal(t, tt.wanted, res.Val)
	}
}
