package l1005

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_largestSumAfterKNegations(t *testing.T) {
	assert.Equal(t, largestSumAfterKNegations([]int{4, 2, 3}, 1), 5)
}
