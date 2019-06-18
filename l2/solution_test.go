package l2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addTwoNumbers(t *testing.T) {
	l1 := createList([]int{2, 4, 3})
	l2 := createList([]int{5, 6, 4})

	res := addTwoNumbers(l1, l2)
	assert.Equal(t, fmt.Sprint(res), "7 0 8")
}

func Test_addCarryNumber(t *testing.T) {
	l1 := createList([]int{5})
	l2 := createList([]int{5})

	res := addTwoNumbers(l1, l2)
	assert.Equal(t, "0 1", fmt.Sprint(res))
}

func Test_addTwoLargeNumbers(t *testing.T) {
	l1 := createList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	l2 := createList([]int{5, 6, 4})

	res := addTwoNumbers(l1, l2)
	assert.Equal(t, fmt.Sprint(res), "6 6 4 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1")
}

func Test_createList(t *testing.T) {
	l := createList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})

	assert.Equal(t, "1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1", fmt.Sprint(l))
}
