package l2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addTwoNumbers(t *testing.T) {
	l1 := numToList(342)
	l2 := numToList(465)

	res := addTwoNumbers(l1, l2)

	assert.Equal(t, res, numToList(807))
}

func Test_addTwoLargeNumbers(t *testing.T) {
	l1 := numToList(1000000000000000000000000000001)
	l2 := numToList(564)

	res := addTwoNumbers(l1, l2)
	assert.Equal(t, res, numToList(6640000000000000000000000000001))
}

func Test_numToList(t *testing.T) {
	l := numToList(342)

	for _, val := range []int{2, 4, 3} {
		assert.Equal(t, val, l.Val)
		l = l.Next
	}
}

func Test_listToNum(t *testing.T) {
	l := numToList(342)

	assert.Equal(t, listToNum(l), uint64(342))
}
