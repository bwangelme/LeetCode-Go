package l525

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMaxLength(t *testing.T) {
	for _, tt := range []struct {
		Arr    []int
		Wanted int
	}{
		{[]int{1, 0, 1, 0, 1, 1, 0, 0}, 8},
		{[]int{1, 0, 1, 0, 1, 1, 0}, 6},
		{[]int{0, 1, 0}, 2},
	} {
		assert.Equal(t, tt.Wanted, findMaxLength(tt.Arr))
	}
}
