package l215

/*
## 题目描述

给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。



示例 1:

输入: [3,2,1,5,6,4], k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4


提示：

1 <= k <= nums.length <= 105
-104 <= nums[i] <= 104

## 解题思路

利用快排的 partition 函数，每次从数组中找到一个数在当前数组的排序索引

然后利用二分法找到合适的值

## 复杂度分析

N = len(nums)

没有申请额外的数组空间，空间复杂度为 O(1)

时间复杂度分析

假设每次找到的数字索引都是在中间

第一次遍历的数组为 N
第二次遍历的数组为 N/2
第三次遍历的数组为 N/4
...
第 logN 次遍历的数组为 1

N+N/2+N/4+...+1 = 2N
*/

func findKthLargestTimeout(nums []int, k int) int {
	var (
		start  = 0
		end    = len(nums) - 1
		target = k - 1
	)
	for {
		idx := partition(nums, start, end)
		if idx == target {
			return nums[idx]
		}

		if idx < target {
			start = idx + 1
		} else {
			end = idx - 1
		}
	}
}

func partition(nums []int, start, end int) int {
	var (
		pivot = nums[end]
		small = start
	)

	for i := start; i < end; i++ {
		if nums[i] > pivot {
			if i != small {
				tmp := nums[i]
				nums[i] = nums[small]
				nums[small] = tmp
			}
			small++
		}
	}

	tmp := nums[end]
	nums[end] = nums[small]
	nums[small] = tmp

	return small
}
