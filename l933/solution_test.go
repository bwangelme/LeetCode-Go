package l933

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RecentCounter(t *testing.T) {
	obj := Constructor()
	assert.Equal(t, 1, obj.Ping(1))
	assert.Equal(t, 2, obj.Ping(100))
	assert.Equal(t, 3, obj.Ping(3001))
	assert.Equal(t, 3, obj.Ping(3002))
}
