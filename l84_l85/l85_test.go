package l84l85

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maximalRectangle(t *testing.T) {
	for _, tt := range []struct {
		Matrix [][]byte
		Wanted int
	}{
		{
			[][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			6,
		},
		{
			[][]byte{
				{'0'},
			},
			0,
		},
		{
			[][]byte{
				{'1'},
			},
			1,
		},
	} {
		assert.Equal(t, tt.Wanted, maximalRectangle(tt.Matrix))
	}
}
