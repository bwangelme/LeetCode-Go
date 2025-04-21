package l540

/*
## 题目描述

给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。

请你找出并返回只出现一次的那个数。

你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。

示例 1:

输入: nums = [1,1,2,3,3,4,4,8,8]
输出: 2
示例 2:

输入: nums =  [3,3,7,7,10,11,11]
输出: 10


提示:

1 <= nums.length <= 105
0 <= nums[i] <= 105

## 解题思路

将数组分成 len(nums) / 2 个二元组, 那么 **第一个不同的二元组，就是唯一数字第一次出现的时刻**

根据二分法进行查找
	1. 如果二元组不等,
		1.1 判断此二元组是否是第一个或前一个二元组是否相等，满足时返回此二元组
		1.2 如果不满足，说明前面还有不等的二元组，将查找范围向左移
	2. 如果二元组相等，说明第一个不等的二元组在后面，将查找范围向右移

## 复杂度分析

N = len(nums)

空间复杂度: O(1), 没有申请新空间
时间复杂度：由于进行的是二分查找，因此时间复杂度是 O(logN/2) = O(logN)
*/

func singleNonDuplicate(nums []int) int {
	var (
		low  = 0
		high = len(nums) / 2
	)

	for low <= high {
		mid := (low + high) / 2
		i := mid * 2

		if i < len(nums)-1 && nums[i] != nums[i+1] {
			if mid == 0 || nums[i-2] == nums[i-1] {
				return nums[i]
			}

			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return nums[len(nums)-1]
}
