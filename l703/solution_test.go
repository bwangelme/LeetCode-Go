package l703

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKthLargest_Add(t *testing.T) {
	kth := Constructor(3, []int{4, 5, 8, 2})
	assert.Equal(t, 4, kth.Add(3))
	assert.Equal(t, 5, kth.Add(5))
	assert.Equal(t, 5, kth.Add(10))
	assert.Equal(t, 8, kth.Add(9))
	assert.Equal(t, 8, kth.Add(4))
}
