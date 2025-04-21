package l35

/*
## 题目描述

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。

示例 1:

输入: nums = [1,3,5,6], target = 5
输出: 2
示例 2:

输入: nums = [1,3,5,6], target = 2
输出: 1
示例 3:

输入: nums = [1,3,5,6], target = 7
输出: 4


提示:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums 为 无重复元素 的 升序 排列数组
-104 <= target <= 104

## 解题思路

利用二分查找查找目标数字

如果目标数字不存在，nums[i-1] < target < nums[i]
那么 i 就是插入位置

边界情况处理

插入位置在数组头部，此时 mid = 0, nums[0] > target, 返回0 即可
插入位置在数组尾部, 此时循环已经结束, high 指向尾部，low = high + 1, nums[n-1] < target, 此时返回 n 即可，插入位置在数组的尾部+1

## 复杂度分析

N == len(nums)

空间复杂度，没有申请新空间，空间复杂度是 O(1)
时间复杂度是 O(logN)

*/

func searchInsert(nums []int, target int) int {
	var (
		low  = 0
		high = len(nums) - 1
	)

	for low <= high {
		mid := (low + high) / 2

		if nums[mid] >= target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}

			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return len(nums)
}
