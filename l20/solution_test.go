package l20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValid(t *testing.T) {
	for _, tt := range []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"(]", false},
		{"]", false},
	} {
		assert.Equal(t, tt.expected, isValid(tt.input))
	}
}
