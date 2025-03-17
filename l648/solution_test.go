package l648

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_replaceWords(t *testing.T) {
	dictionary := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	assert.Equal(t, "the cat was rat by the bat", replaceWords(dictionary, sentence))
}
