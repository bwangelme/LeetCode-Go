package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PrintList(t *testing.T) {
	head := NewList([]string{"1"})
	assert.Equal(t, PrintList(head), "1")

	head = NewList([]string{})
	assert.Equal(t, PrintList(head), "")
}
