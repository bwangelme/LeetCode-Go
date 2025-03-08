package lt

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateList(t *testing.T) {
	l := CreateList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})

	assert.Equal(t, "1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1", fmt.Sprint(l))
}
