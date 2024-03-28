package lcr41

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MovingAverage(t *testing.T) {
	obj := Constructor(3)
	assert.Equal(t, 1.0, obj.Next(1))
	assert.Equal(t, 5.5, obj.Next(10))
	assert.Equal(t, 4.667, obj.Next(3))
	assert.Equal(t, 6.0, obj.Next(5))
}
