package l14

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	for _, tt := range []struct {
		Input  []string
		Wanted string
	}{
		{
			[]string{"flower", "flow", "flight"}, "fl",
		},
		{
			[]string{"dog", "racecar", "car"}, "",
		},
		{
			[]string{""}, "",
		},
		{
			[]string{"dog", "do"}, "do",
		},
	} {
		assert.Equal(t, tt.Wanted, longestCommonPrefix(tt.Input))
	}
}
