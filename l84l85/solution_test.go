package l84l85

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_largestRectangleArea(t *testing.T) {
	for _, tt := range []struct {
		Heights []int
		Wanted  int
	}{
		{
			[]int{2, 1, 5, 6, 2, 3},
			10,
		},
		{
			[]int{2, 4},
			4,
		},
	} {
		//assert.Equal(t, tt.Wanted, largestRectangleArea(tt.Heights))
		assert.Equal(t, tt.Wanted, largestRectangleAreaStack(tt.Heights))
	}
}
