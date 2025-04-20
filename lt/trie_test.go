package lt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Trie(t *testing.T) {
	obj := NewTrie()
	obj.Insert("apple")
	assert.True(t, obj.Search("apple"))
	assert.False(t, obj.Search("app"))
	assert.True(t, obj.StartsWith("app"))
	obj.Insert("app")
	assert.True(t, obj.Search("app"))
}
