package l380

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RandomizedSet(t *testing.T) {
	r := Constructor()
	assert.True(t, r.Insert(1))
	assert.False(t, r.Remove(2))
	assert.True(t, r.Insert(2))
	r.GetRandom()
	assert.True(t, r.Remove(1))
	assert.False(t, r.Insert(2))
	assert.Equal(t, 2, r.GetRandom())
}
