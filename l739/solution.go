package l739

/*
## 题目

给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。

示例 1:

输入: temperatures = [73,74,75,71,69,72,76,73]
输出: [1,1,4,2,1,1,0,0]
示例 2:

输入: temperatures = [30,40,50,60]
输出: [1,1,1,0]
示例 3:

输入: temperatures = [30,60,90]
输出: [1,1,0]

提示：

1 <= temperatures.length <= 10^5
30 <= temperatures[i] <= 100

## 解题思路

将每日温度的索引入栈，入栈前和栈中所有温度比较，找到较高温度的日期出栈

## 复杂度分析

n == len(temperatures)

时间复杂度: O(n), 每日温度入栈出栈的次数只有一次
空间复杂度: O(n), 最坏情况下全都入栈, O(n)
*/

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

func (s *Stack[T]) Empty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Peek() *T {
	stackLen := len(*s)
	if stackLen == 0 {
		return nil
	}
	top := (*s)[stackLen-1]
	return &top
}

func dailyTemperatures(temperatures []int) []int {
	var res = make([]int, len(temperatures))
	s := NewStack[int]()

	for i, _ := range temperatures {
		for !s.Empty() && temperatures[i] > temperatures[*s.Peek()] {
			prev := *s.Pop()
			res[prev] = i - prev
		}

		s.Push(i)
	}

	return res
}
