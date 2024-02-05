package l953

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isAlienSorted(t *testing.T) {
	for _, tt := range []struct {
		Words  []string
		Order  string
		Wanted bool
	}{
		{
			[]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz", true,
		},
		{
			[]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz", false,
		},
	} {
		assert.Equal(t, tt.Wanted, isAlienSorted(tt.Words, tt.Order))
	}
}
