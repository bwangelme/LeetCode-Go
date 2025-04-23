package l875

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minEatingSpeed(t *testing.T) {
	res := minEatingSpeed([]int{312884470}, 968709470)
	assert.Equal(t, 1, res)
}
