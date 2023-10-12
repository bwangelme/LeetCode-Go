package l69

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mySqrt(t *testing.T) {
	for _, tt := range []struct {
		X      int
		Wanted int
	}{
		{4, 2},
		{9, 3},
		{10, 3},
		{8, 2},
		{0, 0},
	} {
		assert.Equal(t, tt.Wanted, mySqrt(tt.X))
	}
}
