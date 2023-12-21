package l234

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPalindromeV2(t *testing.T) {
	for _, tt := range []struct {
		str    string
		wanted bool
	}{
		{"abccba", true},
		{"abcba", true},
		{"1221", true},
		{"12", false},
	} {
		head := NewList(tt.str)
		assert.Equal(t, tt.wanted, isPalindromeV2(head))
	}
}
