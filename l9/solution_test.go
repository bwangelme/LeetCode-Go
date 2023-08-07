package l9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	for _, tt := range []struct {
		Num    int
		Wanted bool
	}{
		{121, true},
		{-121, false},
		{123, false},
	} {
		assert.Equal(t, tt.Wanted, isPalindrome(tt.Num))
		assert.Equal(t, tt.Wanted, isPalindromeNoStr(tt.Num))
	}
}
