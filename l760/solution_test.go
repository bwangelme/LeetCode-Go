package l760

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagramWithSampleCase(t *testing.T) {
	A := []int{12, 28, 46, 32, 50}
	B := []int{50, 12, 32, 46, 28}

	result := anagramMappings(A, B)
	assert.Equal(t, result, []int{1, 4, 3, 2, 0})
}

func TestAnagramWithSameValue(t *testing.T) {
	A := []int{12, 28, 12, 32, 50}
	B := []int{50, 12, 32, 12, 28}

	result := anagramMappings(A, B)
	assert.Equal(t, result, []int{1, 4, 3, 2, 0})
}
