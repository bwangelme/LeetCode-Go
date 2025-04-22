package l528

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Solution(t *testing.T) {
	w := []int{1}
	obj := Constructor(w)
	param_1 := obj.PickIndex()
	assert.Equal(t, param_1, 0)
}
