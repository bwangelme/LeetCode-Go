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
}
