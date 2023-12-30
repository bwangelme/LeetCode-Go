package l209

import "math"

/*
## 题目描述

给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其总和大于等于 target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

## 解题思路

1. 子数组一定是连续的一段
2. 因此可以用 l, r 分别表示子数组的首和尾，我们用最坏 2N 的复杂度遍历完整个数组，求一个最小长度
3. l++ 时，子数组和变小，长度变小
4. r++ 时，子数组和变大，长度变大
5. 利用 3 和 4 相反的性质，水多了加面，面多加水，不断找到合适的值，并求到最小长度

## 复杂度分析

N 表示 len(nums)
时间复杂度

- 尽管有两个循环，但 r, l 在整个循环过程中一直是递增的
- 最坏情况是, nums=[5,3,5,3,5,3], target=8, 此时 l 和 r 都遍历了 2N 遍，时间复杂度依然是 O(N)

空间复杂度

- O(1)

*/

func minSubArrayLen(target int, nums []int) int {
	var (
		sum    = 0
		l      = 0
		r      = 0
		minLen = math.MaxInt
	)

	for r = 0; r < len(nums); r++ {
		sum += nums[r]
		for l <= r && sum >= target {
			minLen = min(minLen, r-l+1)
			sum -= nums[l]
			l++
		}
	}

	if minLen == math.MaxInt {
		return 0
	} else {
		return minLen
	}
}

func min(a, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}
