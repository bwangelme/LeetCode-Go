package l220

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_containsNearbyAlmostDuplicate(t *testing.T) {
	assert.Equal(t, true, containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
	assert.Equal(t, false, containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))
}
