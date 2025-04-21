package l540

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_singleNonDuplicate(t *testing.T) {
	nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	assert.Equal(t, 2, singleNonDuplicate(nums))
	/*
		(1, 1), (2, 3), (3, 4) (4,8), (8)
		n = 9
	*/
}
