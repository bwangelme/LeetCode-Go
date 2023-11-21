package l11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxArea(t *testing.T) {
	assert.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
