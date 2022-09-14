package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircularQueue_NormalOp(t *testing.T) {
	q := NewCircularQueue[string](3)

	assert.True(t, q.Enqueue("apple"))
	assert.True(t, q.Enqueue("pear"))
	assert.True(t, q.Enqueue("peach"))

	assert.True(t, q.IsFull())

	assert.Equal(t, "apple", q.Dequeue())
	assert.Equal(t, "pear", q.Dequeue())
	assert.Equal(t, "peach", q.Dequeue())

	assert.True(t, q.IsEmpty())
}

func TestCircularQueue_EnqueueFull(t *testing.T) {
	q := NewCircularQueue[string](3)

	assert.True(t, q.Enqueue("apple"))
	assert.True(t, q.Enqueue("pear"))
	assert.Equal(t, 2, q.Size())

	assert.True(t, q.Enqueue("peach"))
	assert.False(t, q.Enqueue("banana"))

	assert.True(t, q.IsFull())
}

func TestCircularQueue_Size(t *testing.T) {
	q := NewCircularQueue[string](4)

	assert.True(t, q.Enqueue("Go"))
	assert.True(t, q.Enqueue("Python"))
	assert.True(t, q.Enqueue("Java"))
	assert.True(t, q.Enqueue("PHP"))

	q.Dequeue()
	q.Dequeue()

	assert.True(t, q.Enqueue("JavaScript"))
	assert.True(t, q.Enqueue("Lua"))

	q.Dequeue()

	assert.Equal(t, 3, q.Size())
}
