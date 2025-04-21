package l852

/*
## 题目描述

给定一个长度为 n 的整数 山脉 数组 arr ，其中的值递增到一个 峰值元素 然后递减。

返回峰值元素的下标。

你必须设计并实现时间复杂度为 O(log(n)) 的解决方案。

示例 1：

输入：arr = [0,1,0]
输出：1
示例 2：

输入：arr = [0,2,1,0]
输出：1
示例 3：

输入：arr = [0,10,5,2]
输出：1

提示：

3 <= arr.length <= 105
0 <= arr[i] <= 106
题目数据 保证 arr 是一个山脉数组

## 解题思路

 1. 数组是 >= 3 的，此时数组的最大值必然不可能位于首部或尾部，因此查找范围是 [1, n-2]
 2. 利用二分查找查找数组，
    2.1 如果 nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1], 找到目标值，返回
    2.2 如果 nums[mid] > nums[mid-1], 已经不满足 2.1 的情况下，满足这一点，说明 mid 在上升区间内，向右边移动
    2.3 如果 nums[mid] < nums[mid-1]，在不满足 2.1 的情况下，满足这一点，说明在下降区间内，向左边移动
*/
func peakIndexInMountainArray(arr []int) int {
	var (
		low  int = 1
		high     = len(arr) - 2
	)

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] > arr[mid+1] && arr[mid] > arr[mid-1] {
			return mid
		}

		if arr[mid] < arr[mid+1] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
