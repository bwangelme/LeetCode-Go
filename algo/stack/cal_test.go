package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCal(t *testing.T) {
	for _, tt := range []struct {
		expression string
		wanted     int64
	}{
		{"3+5*8-6", 37},
		{"3+5*8", 43},
		{"30+5*8", 70},
		{"3+2^6", 67},
	} {
		res, err := Cal(tt.expression)
		assert.Nil(t, err)
		assert.Equal(t, tt.wanted, res)
	}
}
