package l55

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canJump(t *testing.T) {
	assert.Equal(t, true, canJump([]int{2, 3, 1, 1, 4}))
	assert.Equal(t, false, canJump([]int{3, 2, 1, 0, 4}))
	assert.Equal(t, true, canJump([]int{2, 0}))
	assert.Equal(t, true, canJump([]int{0}))
}
