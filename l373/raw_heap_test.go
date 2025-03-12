package l373

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_kSmallestPairsV2(t *testing.T) {
	var res [][]int

	res = kSmallestPairsV2([]int{1, 7, 11}, []int{2, 4, 6}, 3)
	assert.Equal(t, res, [][]int{{1, 6}, {1, 2}, {1, 4}})

	res = kSmallestPairsV2([]int{1, 1, 2}, []int{1, 2, 3}, 2)
	assert.Equal(t, res, [][]int{{1, 1}, {1, 1}})
}
