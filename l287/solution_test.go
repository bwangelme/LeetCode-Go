package l287

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findDuplicate(t *testing.T) {
	for _, tt := range []struct {
		Nums   []int
		Wanted int
	}{
		{
			[]int{1, 3, 4, 2, 2},
			2,
		},
		{
			[]int{3, 1, 3, 4, 2},
			3,
		},
	} {
		assert.Equal(t, tt.Wanted, findDuplicate(tt.Nums))
	}
}
