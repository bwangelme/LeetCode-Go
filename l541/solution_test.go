package l541

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse(t *testing.T) {
	assert.Equal(t, reverse("abcd"), "dcba")
	assert.Equal(t, reverse("上海自来水来自海上"), "上海自来水来自海上")
}

func Test_reverserStr(t *testing.T) {
	assert.Equal(t, "bacdfeg", reverseStr("abcdefg", 2))
	assert.Equal(t, "bacd", reverseStr("abcd", 2))
}
