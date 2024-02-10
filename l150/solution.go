package l150

import "strconv"

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

/*
## 题目描述

给你一个字符串数组 tokens ，表示一个根据 逆波兰表示法 表示的算术表达式。

请你计算该表达式。返回一个表示表达式值的整数。

注意：

有效的算符为 '+'、'-'、'*' 和 '/' 。
每个操作数（运算对象）都可以是一个整数或者另一个表达式。
两个整数之间的除法总是 向零截断 。
表达式中不含除零运算。
输入是一个根据逆波兰表示法表示的算术表达式。
答案及所有中间计算结果可以用 32 位 整数表示。

## 复杂度分析

n == len(tokens)
时间复杂度: O(n),
遇到运算符, pop 两次，计算一次，push 一次，遇到 num, push 一次。总的运算次数比 n 多，但趋势还是 O(n)

空间复杂度: O(n)
*/
func evalRPN(tokens []string) int {
	s := NewStack[int]()
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			num2 := s.Pop()
			num1 := s.Pop()
			s.Push(calculate(*num1, *num2, token))
		default:
			v, _ := strconv.Atoi(token)
			s.Push(v)
		}
	}

	return *s.Pop()
}

func calculate(num1, num2 int, op string) int {
	switch op {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	}

	return 0
}
