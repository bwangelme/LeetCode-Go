package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppend(t *testing.T) {
	for _, tt := range []struct {
		Head    *SingleListNode[string]
		Val     string
		Wanted1 string
		Wanted2 string
	}{
		{
			Head:    NewList([]string{"a", "b", "c"}),
			Val:     "end",
			Wanted1: "a -> b -> c -> end",
			Wanted2: "a -> b -> c -> end -> end",
		},
	} {
		var head *SingleListNode[string]
		head = Append(tt.Head, tt.Val)
		assert.Equal(t, tt.Wanted1, PrintList(head))
		head = AppendDummy(tt.Head, tt.Val)
		assert.Equal(t, tt.Wanted2, PrintList(head))
	}
}

func TestDeleteVal(t *testing.T) {
	for _, tt := range []struct {
		Head   *SingleListNode[string]
		Val    string
		Wanted string
	}{
		{
			Head:   NewList([]string{"a", "delete", "b", "c"}),
			Val:    "delete",
			Wanted: "a -> b -> c",
		},
	} {
		head := DeleteValDummy(tt.Head, tt.Val)
		//head := DeleteVal(tt.Head, tt.Val)
		assert.Equal(t, tt.Wanted, PrintList(head))
	}
}
