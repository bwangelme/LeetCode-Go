package eight_queen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetQueen(t *testing.T) {
	res := GetQueen(8)
	assert.Equal(t, res, 92)
}
