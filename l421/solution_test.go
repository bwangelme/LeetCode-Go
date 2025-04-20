package l421

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMaximumXOR(t *testing.T) {
	max := findMaximumXOR([]int{3, 10, 5, 25, 2, 8})
	assert.Equal(t, max, 28)
}
