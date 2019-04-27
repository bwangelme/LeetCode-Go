package l231

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPowerOfTwo(t *testing.T) {
	assert.Equal(t, true, isPowerOfTwo(1))
	assert.Equal(t, true, isPowerOfTwo(16))
	assert.Equal(t, false, isPowerOfTwo(218))
}
