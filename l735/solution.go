package l735

/*
## 题目

给定一个整数数组 asteroids，表示在同一行的小行星。

对于数组中的每一个元素，其绝对值表示小行星的大小，正负表示小行星的移动方向（正表示向右移动，负表示向左移动）。每一颗小行星以相同的速度移动。

找出碰撞后剩下的所有小行星。碰撞规则：两个小行星相互碰撞，较小的小行星会爆炸。如果两颗小行星大小相同，则两颗小行星都会爆炸。两颗移动方向相同的小行星，永远不会发生碰撞。

示例 1：

输入：asteroids = [5,10,-5]
输出：[5,10]
解释：10 和 -5 碰撞后只剩下 10 。 5 和 10 永远不会发生碰撞。
示例 2：

输入：asteroids = [8,-8]
输出：[]
解释：8 和 -8 碰撞后，两者都发生爆炸。
示例 3：

输入：asteroids = [10,2,-5]
输出：[10]
解释：2 和 -5 发生碰撞后剩下 -5 。10 和 -5 发生碰撞后剩下 10 。

提示：

2 <= asteroids.length <= 10^4
-1000 <= asteroids[i] <= 1000
asteroids[i] != 0

## 解题思路

根据题目描述，可得，

- 向左飞时直接入栈
- 向右飞时，碰撞栈内每一颗左飞的星星

思路

1. 向右飞的行星 as 比栈中所有左飞行星都大，不断弹栈
2. 向右飞的行星 as 和栈顶的左飞行星相等，相互消失
3. 其他情况, 直接入栈
3.1 as 是左飞行星
3.2 右飞行星撞穿了栈，栈为空
3.3 右飞行星撞穿了栈，栈里只有右飞行星了
4. 将栈中所有行星按照从底到顶的顺序返回

## 复杂度分析

n == len(asteroids)

由于每个行星只会入栈出栈一次， 时间复杂度 O(n)

空间复杂度为 O(n), 使用的额外空间就是栈的空间，最坏为 O(n)
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

func (s *Stack[T]) ToArray() []T {
	return *s
}

func asteroidCollision(asteroids []int) []int {
	var s = NewStack[int]()
	for _, as := range asteroids {
		for !s.Empty() && *s.Peek() > 0 && *s.Peek() < -as {
			s.Pop()
		}

		if !s.Empty() && as < 0 && *s.Peek() == -as {
			s.Pop()
		} else if as > 0 || s.Empty() || *s.Peek() < 0 {
			s.Push(as)
		}
	}

	return s.ToArray()
}
