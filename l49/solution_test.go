package l49

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	for _, tt := range []struct {
		Input  []string
		Wanted [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
	} {
		assert.Equal(t, tt.Wanted, groupAnagrams(tt.Input))
	}
}
