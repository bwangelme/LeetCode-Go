package l242

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	for _, tt := range []struct {
		S      string
		T      string
		Wanted bool
	}{
		{"孙行者", "行者孙", true},
		{"rat", "car", false},
		{"anagram", "nagaram", true},
	} {
		assert.Equal(t, tt.Wanted, isAnagram(tt.S, tt.T))
	}
}
