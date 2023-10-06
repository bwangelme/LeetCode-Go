package l844

/**
## 题目

给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，如果两者相等，返回 true 。# 代表退格字符。

注意：如果对空文本输入退格字符，文本继续为空。

## 复杂度分析

### 时间复杂度

往栈中写入数据次数和 s, t 这两个字符串长度有关。

所以时间复杂度为 O(m+n), m == len(s), n == len(t)

### 空间复杂度

栈的大小和字符串的长度有关，所以空间复杂度为 O(m+n)
*/

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
	return &top
}

func (s *Stack[T]) String() string {
	if len(*s) == 0 {
		return ""
	}

	var b strings.Builder
	for i := 0; i < len(*s); i++ {
		b.WriteString(fmt.Sprintf("%v", (*s)[i]))
	}

	return b.String()
}

func backspaceCompare(s string, t string) bool {
	var (
		sStack = NewStack[rune]()
		tStack = NewStack[rune]()
	)
	for _, tt := range []struct {
		str   string
		stack *Stack[rune]
	}{
		{s, &sStack},
		{t, &tStack},
	} {
		for _, c := range tt.str {
			if c == '#' {
				tt.stack.Pop()
				continue
			}
			tt.stack.Push(c)
		}
	}

	return sStack.String() == tStack.String()
}
