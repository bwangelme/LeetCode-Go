package l84l85

/*
## 题目

给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。

提示：

1 <= heights.length <=1e5
0 <= heights[i] <= 1e4
*/

/*
## 解题思路

分治法

heights 中最低的柱子的索引是 minHeightIndex

面积最大的数组可以分为三种情况

1. 穿过 minHeightIndex 柱子, 此时高度最低，求最大需要让长度最长，因此长度是整个数组
2. 在 minHeightIndex 的左边
3. 在 minHeightIndex 的右边

利用递归依次求出三种情况, 然后求出最大值

## 复杂度分析

n == len(height)

最坏情况

minHeightsIndex 每次都在数组最左边，此时递归深度是 O(n), 分治过程是一个链表

时间复杂度: O(n^2)
空间复杂度: O(n) // helper 函数的递归深度

最好情况

minHeightsIndex 在数组中间，此时递归深度是 O(logN), 分治过程是一颗二叉树

时间复杂度: O(NlogN)
空间复杂度: O(logN) // helper 函数的递归深度
*/
func largestRectangleArea(heights []int) int {
	n := len(heights)
	return helper(heights, 0, n)
}

func helper(heights []int, start, end int) int {
	// 一根柱子也没有
	if start == end {
		return 0
	}

	// 只有 start 一根柱子
	if start+1 == end {
		return heights[start]
	}

	minHeightIndex := start
	for i := start; i < end; i++ {
		if heights[i] < heights[minHeightIndex] {
			minHeightIndex = i
		}
	}

	//[]int{2, 1, 5, 6, 2, 3},
	area := (end - start) * heights[minHeightIndex]
	left := helper(heights, start, minHeightIndex)
	right := helper(heights, minHeightIndex+1, end)

	area = max(area, left)
	return max(area, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
