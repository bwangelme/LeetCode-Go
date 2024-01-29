package l125

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	for _, tt := range []struct {
		Input  string
		Wanted bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
	} {
		assert.Equal(t, tt.Wanted, isPalindrome(tt.Input))
	}
}
