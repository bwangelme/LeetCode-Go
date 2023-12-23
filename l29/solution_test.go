package l29

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivide(t *testing.T) {
	assert.Equal(t, divide(-9, -3), 3)
	assert.Equal(t, divide(10, 3), 3)
	assert.Equal(t, divide(7, -3), -2)
	assert.Equal(t, divide(-2147483648, -1), 2147483647)
	assert.Equal(t, 1073741824, divide(-2147483648, -2))
}

func TestDivideV2(t *testing.T) {
	assert.Equal(t, 3, divideV2(-9, -3))
	assert.Equal(t, 3, divideV2(10, 3))
	assert.Equal(t, -2, divideV2(7, -3))
	assert.Equal(t, 2147483647, divideV2(-2147483648, -1))
	assert.Equal(t, 1073741824, divideV2(-2147483648, -2))
}
