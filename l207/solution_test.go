package l207

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canFinish(t *testing.T) {
	assert.Equal(t, true, canFinish(2, [][]int{{1, 0}}))
	assert.Equal(t, false, canFinish(2, [][]int{{1, 0}, {0, 1}}))
}
