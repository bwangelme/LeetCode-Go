package l318

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxProduct(t *testing.T) {
	for _, tt := range []struct {
		Input  []string
		Wanted int
	}{
		{
			[]string{"abcw", "foo", "bar", "fxyz", "abcdef"},
			16,
		},
	} {
		assert.Equal(t, tt.Wanted, maxProduct(tt.Input))
	}
}
