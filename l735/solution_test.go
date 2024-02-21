package l735

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_asteroidCollision(t *testing.T) {
	for _, tt := range []struct {
		Input  []int
		Wanted []int
	}{
		{
			[]int{5, 10, -5}, []int{5, 10},
		},
		{
			[]int{8, -8}, []int{},
		},
		{
			[]int{10, 2, -5}, []int{10},
		},
	} {
		assert.Equal(t, tt.Wanted, asteroidCollision(tt.Input))
	}

}
