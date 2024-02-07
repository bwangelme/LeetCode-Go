package l539

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMinDifference(t *testing.T) {
	for _, tt := range []struct {
		Input  []string
		Wanted int
	}{
		{[]string{"23:59", "00:00"}, 1},
		{[]string{"00:00", "23:59", "00:00"}, 0},
	} {
		assert.Equal(t, tt.Wanted, findMinDifference(tt.Input))
	}
}
