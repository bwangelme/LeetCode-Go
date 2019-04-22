package l338

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBits(t *testing.T) {
	assert.Equal(t, []int{0, 1, 1}, []int{0, 1, 1})
	assert.Equal(t, CountBits(2), []int{0, 1, 1})
	assert.Equal(t, CountBits(5), []int{0, 1, 1, 2, 1, 2})
}
