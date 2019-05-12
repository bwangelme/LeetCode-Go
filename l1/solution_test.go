package l1

import (
	"github.com/d5/tengo/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	input := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(input, target)
	assert.Equal(t, result, []int{0, 1})
}
