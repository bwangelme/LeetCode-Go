package l151

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	assert.Equal(t, reverseWords("a good   example"), "example good a")
	assert.Equal(t, reverseWords(""), "")
	assert.Equal(t, reverseWords(" "), "")
	assert.Equal(t, reverseWords("the sky is blue"), "blue is sky the")
	assert.Equal(t, reverseWords("  hello world!  "), "world! hello")
}
