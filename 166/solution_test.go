package l166

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHammingWeight(t *testing.T) {
	assert.Equal(t, hammingWeight(binStrToUint32("00000000000000000000000000001011")), 3)
	assert.Equal(t, hammingWeight(binStrToUint32("11111111111111111111111111111101")), 31)
	assert.Equal(t, hammingWeight(binStrToUint32("00000000000000000000000010000000")), 1)
	assert.Equal(t, hammingWeight(binStrToUint32("00000000000000000000000000000000")), 0)
}
