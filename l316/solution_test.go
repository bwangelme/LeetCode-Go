package l316

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeDuplicateLetters(t *testing.T) {
	assert.Equal(t, "acdb", removeDuplicateLetters("cbacdcbc"))
	assert.Equal(t, "abc", removeDuplicateLetters("bcabc"))
}
