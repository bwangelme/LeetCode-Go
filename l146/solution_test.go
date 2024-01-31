package l146

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LRUCache(t *testing.T) {
	l := Constructor(2)
	l.Put(1, 1) // 缓存是 {1=1}
	l.Put(2, 2) // 缓存是 {1=1, 2=2}
	assert.Equal(t, 1, l.Get(1))
	l.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	assert.Equal(t, -1, l.Get(2))
	l.Put(4, 4)                   // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	assert.Equal(t, -1, l.Get(1)) // 返回 -1 (未找到)
	assert.Equal(t, 3, l.Get(3))  // 返回 3
	assert.Equal(t, 4, l.Get(4))  // 返回 4

}
