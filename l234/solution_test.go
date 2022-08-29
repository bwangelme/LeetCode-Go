package l234

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewList(t *testing.T) {
	l := NewList("abccba")
	res := PrintList(l)
	assert.Equal(t, "a -> b -> c -> c -> b -> a", res)
}

func Test_isPalindrome(t *testing.T) {
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
		assert.Equal(t, tt.wanted, isPalindrome(head))
	}
}
