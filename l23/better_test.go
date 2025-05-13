package l23

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeKListsBetter(t *testing.T) {
	var lists = []*ListNode{
		makeList([]int{1, 4, 5}),
		makeList([]int{1, 3, 4}),
		makeList([]int{2, 6}),
	}

	res := mergeKListsBetter(lists)
	assert.True(t, listEqual(makeList([]int{1, 1, 2, 3, 4, 4, 5, 6}), res))

	lists = []*ListNode{
		makeList([]int{2}),
		nil,
		makeList([]int{-1}),
	}

	res = mergeKListsBetter(lists)
	assert.True(t, listEqual(makeList([]int{-1, 2}), res))
}
