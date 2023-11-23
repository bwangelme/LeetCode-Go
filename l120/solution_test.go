package l120

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minimumTotal(t *testing.T) {
	triangle := [][]int{
		{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3},
	}
	assert.Equal(t, 11, minimumTotal(triangle))

	triangle = [][]int{
		{-10},
	}
	assert.Equal(t, -10, minimumTotal(triangle))
}

func Test_minimumTotalV2(t *testing.T) {
	triangle := [][]int{
		{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3},
	}
	assert.Equal(t, 11, minimumTotalV2(triangle))

	triangle = [][]int{
		{-10},
	}
	assert.Equal(t, -10, minimumTotalV2(triangle))
}
