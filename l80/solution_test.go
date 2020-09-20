package l80

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RemoveDuplicates(t *testing.T) {
	assert.Equal(t, 5, removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
	assert.Equal(t, 7, removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
}
