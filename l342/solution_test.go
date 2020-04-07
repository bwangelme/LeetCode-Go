package l342

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfFour(t *testing.T) {
	assert.True(t, isPowerOfFour(4))
	assert.True(t, isPowerOfFour(1))
	assert.False(t, isPowerOfFour(5))
}
