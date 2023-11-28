package l62

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_uniquePaths(t *testing.T) {
	assert.Equal(t, 3, uniquePaths(3, 2))
}
