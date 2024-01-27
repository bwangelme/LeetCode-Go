package l76

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minWindow(t *testing.T) {
	for _, tt := range []struct {
		S      string
		T      string
		Wanted string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"abcdaeegf", "aab", "abcda"},
	} {
		assert.Equal(t, tt.Wanted, minWindow(tt.S, tt.T))
	}
}
