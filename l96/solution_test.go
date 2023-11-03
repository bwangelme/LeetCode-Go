package l96

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numTrees(t *testing.T) {
	assert.Equal(t, 5, numTrees(3))
}
