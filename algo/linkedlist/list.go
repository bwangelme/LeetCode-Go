package linkedlist

import (
	"fmt"
	"strings"
)

//SingleListNode 单链表
type SingleListNode[T any] struct {
	Data T
	Next *SingleListNode[T]
}

func (s SingleListNode[T]) String() string {
	return fmt.Sprintf("%v", s.Data)
}

func NewSingleListNode[T any](data T, next *SingleListNode[T]) *SingleListNode[T] {
	return &SingleListNode[T]{
		Data: data,
		Next: next,
	}
}

func PrintList[T any](head *SingleListNode[T]) string {
	var b = strings.Builder{}
	var current = head

	for {
		if current == nil {
			break
		}

		b.WriteString(fmt.Sprintf("%v", current.Data))
		if current.Next != nil {
			b.WriteString(" -> ")
		}
		current = current.Next
	}

	return b.String()
}

func NewList[T any](arr []T) *SingleListNode[T] {
	var (
		head    *SingleListNode[T] = nil
		current                    = head
	)

	for _, r := range arr {
		if head == nil {
			head = NewSingleListNode(r, nil)
			current = head
		} else {
			current.Next = NewSingleListNode(r, nil)
			current = current.Next
		}
	}

	return head
}
