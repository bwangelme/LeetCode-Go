package l11

/*
## 题目

给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。

n == height.length
2 <= n <= 10^5
0 <= height[i] <= 10^4

## 解题思路

假设两条边的索引分别是 l, r

容器的面积 area = (r-l) * min(height[l], height[r])

当我们向内移动边的时候，(r-l) 一定会变小

当我们移动长边的时候，min(height[l], height[r]), 可能不变或者变小
当我们移动短边的时候，min(height[l], height[r]) 可能变大，变小, 不变

为了求最大的面积，我们需要每次移动短边，因此解题思路就是，从 [0, n-1] 开始，不断移动短边计算面积，直到 l, r 相遇。此过程中可以计算出最大面积。

## 复杂度分析

n 表示数组的长度

### 时间复杂度

时间复杂度是 O(n)

### 空间复杂度

没有申请和 n 相关的额外的数组空间，空间复杂度是 O(1)

*/

func maxArea(height []int) int {
	var (
		max   = 0
		left  = 0
		right = len(height) - 1
		area  int
	)
	for left < right {
		d := right - left
		if height[left] < height[right] {
			area = d * height[left]
			left++
		} else {
			area = d * height[right]
			right--
		}
		if area > max {
			max = area
		}
	}

	return max
}
