package l680

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_validPalindrome(t *testing.T) {
	for _, tt := range []struct {
		Input  string
		Wanted bool
	}{
		{"aba", true},
		{"abca", true},
		{"abc", false},
	} {
		assert.Equal(t, tt.Wanted, validPalindrome(tt.Input))
	}
}
