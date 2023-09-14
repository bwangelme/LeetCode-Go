package find

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Sqrt(t *testing.T) {
	for _, tt := range []struct {
		Input  float64
		Wanted float64
	}{
		{
			16,
			4,
		},
	} {
		assert.Equal(t, tt.Wanted, Sqrt(tt.Input))
	}
}
