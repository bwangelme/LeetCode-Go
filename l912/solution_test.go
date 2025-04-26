package l912

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortArray(t *testing.T) {
	nums := []int{5, 2, 3, 1}
	assert.Equal(t, sortArray(nums), []int{1, 2, 3, 5})

	nums = []int{110, 100, 0}
	assert.Equal(t, sortArray(nums), []int{0, 100, 110})
}
