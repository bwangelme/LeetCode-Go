package l84l85

import "github.com/bwangelme/LeetCode-Go/lt"

/*
直方图最大矩形面积，单调栈解法
===

## 解题思路

直方图中，以 **某个柱子为顶** 的 **最大面积** ，是这个柱子向两边不断扩展，直到遇到了比他矮的柱子。

栈中存储的柱子是单调递增的。栈中元素的前一位，就是比栈顶柱子矮的前一个柱子

如果栈中只有一个元素，表示这个柱子当前是最矮的，左边没有比它更矮的柱子了，可以使用 i - -1 -1 来计算宽度，这包括了当前栈顶柱子左边的所有柱子

如果遍历完了 heights 数组，栈中依然有元素，表示栈中的柱子比后续的柱子都矮，右边就可以计算它们到数组末尾的索引。

## 复杂度分析

n == len(heights)

每个柱子都只会入栈出栈一次，出栈后计算面积的逻辑也只会执行一次，因此时间复杂度是 O(n)

借用了一个栈作为辅助空间，因此空间复杂度最坏是 O(n)
*/
func largestRectangleAreaStack(heights []int) int {
	s := lt.NewStack[int]()
	s.Push(-1)

	var maxArea = 0
	for i := 0; i < len(heights); i++ {
		for *s.Top() != -1 && heights[*s.Top()] >= heights[i] {
			height := heights[*s.Pop()]
			width := i - *s.Top() - 1
			maxArea = max(maxArea, width*height)
		}

		s.Push(i)
	}

	for *s.Top() != -1 {
		height := heights[*s.Pop()]
		width := len(heights) - *s.Top() - 1
		maxArea = max(maxArea, width*height)
	}

	return maxArea
}
