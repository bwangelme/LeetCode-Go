package l347

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	res := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	assert.Equal(t, []int{2, 1}, res)
}
