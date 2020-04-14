package l343

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_integerBreak(t *testing.T) {
	for _, tt := range []struct {
		number int
		want   int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 4},
		{5, 6},
		{10, 36},
	} {
		assert.Equal(t, tt.want, integerBreak(tt.number))
	}
}
