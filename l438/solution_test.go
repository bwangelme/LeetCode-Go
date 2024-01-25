package l438

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	for _, tt := range []struct {
		S      string
		P      string
		Wanted []int
	}{
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abab", "ab", []int{0, 1, 2}},
	} {
		assert.Equal(t, tt.Wanted, findAnagrams(tt.S, tt.P))
	}
}
