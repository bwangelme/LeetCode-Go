package l88

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_merge(t *testing.T) {
	var (
		nums1 = []int{1, 2, 3, 0, 0, 0}
		m     = 3
		nums2 = []int{2, 5, 6}
		n     = 3
	)
	merge(nums1, m, nums2, n)
	assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, nums1)
}

func Test_merge_empty(t *testing.T) {
	var (
		nums1 = []int{1}
		m     = 1
		nums2 = []int{}
		n     = 0
	)
	merge(nums1, m, nums2, n)
	assert.Equal(t, []int{1}, nums1)
}
