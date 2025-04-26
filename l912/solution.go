package l912

import "math/rand"

/*
## 题目描述

给你一个整数数组 nums，请你将该数组升序排列。

你必须在 不使用任何内置函数 的情况下解决问题，时间复杂度为 O(nlog(n))，并且空间复杂度尽可能小。



示例 1：

输入：nums = [5,2,3,1]
输出：[1,2,3,5]
示例 2：

输入：nums = [5,1,1,2,0,0]
输出：[0,0,1,1,2,5]


提示：

1 <= nums.length <= 5 * 104
-5 * 104 <= nums[i] <= 5 * 104

## 解题思路

实现快排

## 复杂度分析

N == len(nums)

取决于 random 取到的 p 的值在数组中大小的顺序
	如果每次取到的是最大值，时间复杂度是 O(n^2)
	如果每次取到的是中间值，时间复杂度是 O(nLogN)
*/

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	if end > start {
		pivot := partition(nums, start, end)
		quickSort(nums, start, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}

func partition(nums []int, start, end int) int {
	p := rand.Intn(end-start+1) + start
	swap(nums, p, end)
	var small = start - 1

	for i := start; i < end; i++ {
		if nums[i] < nums[end] {
			small++
			swap(nums, small, i)
		}
	}

	small++
	swap(nums, small, end)

	return small
}

func swap(nums []int, idx1, idx2 int) {
	if idx1 != idx2 {
		temp := nums[idx1]
		nums[idx1] = nums[idx2]
		nums[idx2] = temp
	}
}
