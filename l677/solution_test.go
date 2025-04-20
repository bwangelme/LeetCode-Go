package l677

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MapSum(t *testing.T) {
	obj := Constructor()
	obj.Insert("apple", 3)
	param_2 := obj.Sum("ap")
	assert.Equal(t, param_2, 3)
}
