package l77

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combine(t *testing.T) {
	result := combine(4, 2)
	assert.Equal(t, result, [][]int{
		[]int{3, 4}, []int{2, 4}, []int{2, 3}, []int{1, 4}, []int{1, 3}, []int{1, 2},
	})
}
