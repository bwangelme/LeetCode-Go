package l208

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_UTrie(t *testing.T) {
	obj := NewUTrie()
	obj.Insert("飞流直下三千尺")
	assert.True(t, obj.Search("飞流直下三千尺"))
	assert.False(t, obj.Search("飞流"))
	assert.True(t, obj.StartsWith("飞流"))
	obj.Insert("飞流")
	assert.True(t, obj.Search("飞流"))
}
