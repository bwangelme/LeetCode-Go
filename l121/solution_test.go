package l121

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	assert.Equal(t, 5, maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
