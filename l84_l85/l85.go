package l84l85

/*
## 题目描述

给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。

提示：

rows == matrix.length
cols == matrix[0].length
1 <= row, cols <= 200
matrix[i][j] 为 '0' 或 '1'

## 解题思路

将矩阵中的包含 1 的图形，转换成 84 题中的直方图，就可以用84题的解法来解了

以矩阵中每一行为基线，将矩阵中包含1的各自转换成直方图

## 复杂度分析

matrix 为 m 行 n 列

时间复杂度:

针对每一行，计算 maxArea 的时间复杂度为 O(n)
总体时间复杂度为 O(m * n)

空间复杂度:

使用了辅助数组 heigths, 空间复杂度为 O(n)
*/

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	var maxArea = 0
	var heights = make([]int, len(matrix[0]))
	for _, row := range matrix {
		for i, c := range row {
			if c == '0' {
				heights[i] = 0
			} else {
				heights[i]++
			}
		}
		maxArea = max(maxArea, largestRectangleAreaStack(heights))
	}

	return maxArea
}
