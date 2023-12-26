package l338

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBits(t *testing.T) {
	for _, tt := range []struct {
		Input  int
		Wanted []int
	}{
		{
			2, []int{0, 1, 1},
		},
		{
			5, []int{0, 1, 1, 2, 1, 2},
		},
	} {
		assert.Equal(t, tt.Wanted, CountBits(tt.Input))
		assert.Equal(t, tt.Wanted, CountBitsV2(tt.Input))
	}
}
