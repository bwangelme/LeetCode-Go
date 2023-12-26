package l137

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	for _, tt := range []struct {
		Input  []int
		Wanted int
	}{
		{[]int{2, 2, 3, 2}, 3},
		{[]int{0, 1, 0, 1, 0, 1, 99}, 99},
		{[]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}, -4},
	} {
		assert.Equal(t, tt.Wanted, singleNumber(tt.Input))
	}
}
