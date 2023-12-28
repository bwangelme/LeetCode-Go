package l167

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_twoSum(t *testing.T) {
	for _, tt := range []struct {
		Input  []int
		Target int
		Wanted []int
	}{
		{
			[]int{2, 3, 4},
			6,
			[]int{1, 3},
		},
	} {
		assert.Equal(t, tt.Wanted, twoSum(tt.Input, tt.Target))
	}
}
