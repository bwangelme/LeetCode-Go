package l647

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countSubstrings(t *testing.T) {
	for _, tt := range []struct {
		Input  string
		Wanted int
	}{
		{"abc", 3},
		{"aaa", 6},
	} {
		assert.Equal(t, tt.Wanted, countSubstrings(tt.Input))
	}
}
