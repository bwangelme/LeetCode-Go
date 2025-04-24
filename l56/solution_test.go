package l56

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_merge(t *testing.T) {
	intervals := [][]int{
		{1, 4},
		{4, 5},
	}
	res := merge(intervals)
	assert.Equal(t, res, [][]int{{1, 5}})
}
