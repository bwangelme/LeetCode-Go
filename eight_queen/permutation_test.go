package eight_queen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Permutation(t *testing.T) {
	results := Permutation(3)
	assert.Equal(t, 6, len(results))

	assert.Equal(t, []Per{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}, results)

	results = Permutation(8)
	assert.Equal(t, len(results), 40320) // A^8_8 = 8! / 0! = 40320
}

func Test_tmp(t *testing.T) {
	tmp()
}
