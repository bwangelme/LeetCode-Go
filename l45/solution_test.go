package l45

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_jump(t *testing.T) {
	assert.Equal(t, 2, jump([]int{2, 3, 1, 1, 4}))
	assert.Equal(t, 0, jump([]int{0}))
	assert.Equal(t, 2, jump([]int{2, 3, 0, 1, 4}))
}
