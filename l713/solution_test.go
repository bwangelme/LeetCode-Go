package l713

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numSubarrayProductLessThanK(t *testing.T) {
	assert.Equal(t, 8, numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
}
