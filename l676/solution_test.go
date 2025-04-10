package l676

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MagicDictionary(t *testing.T) {
	obj := Constructor()
	dictionary := []string{"hello", "leetcode"}

	obj.BuildDict(dictionary)
	assert.Equal(t, false, obj.Search("hello"))
	assert.Equal(t, true, obj.Search("hhllo"))
	assert.Equal(t, true, obj.Search("aello"))

}
