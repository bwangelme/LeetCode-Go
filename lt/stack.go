package lt

import (
	"fmt"
	"strings"
)

type Stack[T any] []T

func NewStack[T any]() Stack[T] {
	return Stack[T]{}
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) Push(elem T) bool {
	*s = append(*s, elem)

	return true
}

func (s *Stack[T]) Pop() *T {
	stackLen := len(*s)
	if stackLen == 0 {
		return nil
	}

	// 注意: *s 必须扩起来，[] 的优先级高于 *
	top := (*s)[stackLen-1]
	*s = (*s)[:stackLen-1]

	return &top
}

func (s *Stack[T]) Top() *T {
	stackLen := len(*s)
	if stackLen == 0 {
		return nil
	}

	top := (*s)[stackLen-1]
	// TODO: 外部可以修改栈内的内容，这里应该返回值而不是指针
	return &top
}

func (s *Stack[T]) String() string {
	if len(*s) == 0 {
		return "Empty"
	}

	var b strings.Builder
	for i := len(*s) - 1; i >= 0; i-- {
		b.WriteString(fmt.Sprintf("%d %v\n", i, (*s)[i]))
	}

	return b.String()
}
