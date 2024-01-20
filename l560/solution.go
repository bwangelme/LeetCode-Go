package l560

/*
## 题目描述

给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。

子数组是数组中元素的连续非空序列。



示例 1：

输入：nums = [1,1,1], k = 2
输出：2
示例 2：

输入：nums = [1,2,3], k = 3
输出：2


提示：

1 <= nums.length <= 2 * 10^4
-1000 <= nums[i] <= 1000
-10^7 <= k <= 10^7
*/

func subarraySum(nums []int, k int) int {
	var sumToCount = make(map[int]int, 0)
	sumToCount[0] = 1

	sum := 0
	count := 0
	for _, num := range nums {
		sum += num
		v, ok := sumToCount[sum-k]
		if ok {
			count += v
		}

		v, ok = sumToCount[sum]
		if ok {
			sumToCount[sum] = v + 1
		} else {
			sumToCount[sum] = 1
		}
	}

	return count
}
