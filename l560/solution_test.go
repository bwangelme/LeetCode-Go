package l560

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_subarraySum(t *testing.T) {
	assert.Equal(t, 2, subarraySum([]int{1, 1, 1}, 2))
	assert.Equal(t, 2, subarraySum([]int{1, 2, 3}, 3))
}
