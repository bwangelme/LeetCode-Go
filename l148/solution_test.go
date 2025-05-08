package l148

import (
	"fmt"
	"testing"

	"github.com/bwangelme/LeetCode-Go/lt"
	"github.com/stretchr/testify/assert"
)

func Test_sortList(t *testing.T) {
	l := lt.CreateList([]int{4, 2, 1, 3})
	l = sortList(l)
	v := fmt.Sprint(l)
	assert.Equal(t, v, "1 2 3 4")
}
