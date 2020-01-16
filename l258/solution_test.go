package l258

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddDigits(t *testing.T) {
	assert.Equal(t, 2, addDigits(38))
	assert.Equal(t, 0, addDigits(0))
	assert.Equal(t, 9, addDigits(9))
	assert.Equal(t, 9, addDigits(18))
}
