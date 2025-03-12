package lt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	h := NewHeap[int](func(v1 int, v2 int) bool {
		return v1 > v2
	})
	for i := 20; i > 0; i-- {
		h.Offer(i) // all elements are the same
	}

	total := h.Len()
	for i := 1; h.Len() > 0; i++ {
		x := h.Poll()
		assert.Equal(t, x, total+1-i)
	}
	assert.Equal(t, 0, h.Len())
}
