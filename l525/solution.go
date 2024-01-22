package l525

/*
## 题目

给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。

示例 1:

输入: nums = [0,1]
输出: 2
说明: [0, 1] 是具有相同数量 0 和 1 的最长连续子数组。
示例 2:

输入: nums = [0,1,0]
输出: 2
说明: [0, 1] (或 [1, 0]) 是具有相同数量0和1的最长连续子数组。


提示：

1 <= nums.length <= 10^5
nums[i] 不是 0 就是 1

## 解题思路

1. 将数组中的 0 全部转换成 -1，那么就可以将问题转换成，和为 0 的最长连续子数组的长度
2. S[i] 表示 nums[0,i] 所有数字的和，用一个 map 保存 S[i] 和第一次遇到S[i]时 i 的值
3. 求和为 0 的连续子数组，就是求两个相同的 S[i] 的差值
4. S[i] == 0 是一个特殊值，如果 S[i] == 0, 那么整个数组就是我们要的结果，这个数组的长度就是 i+1 == i - (-1), 因此我们设置 sumToIndex[0] = 1

*/

func findMaxLength(nums []int) int {
	// S[i] 表示 nums[0,i] 所有数字的和
	// sumToIndex 的键是 S[i], 值是第一次遇到 S[i] 时 i 的索引
	sumToIndex := make(map[int]int)

	// 如果 S[i] == 0, 那么表示 nums[0,i] 就是符合条件的数组, len(nums[0,i]) == i+1 == i - (-1)
	// 因此 sumToIndex[0] = -1
	sumToIndex[0] = -1

	var (
		sum       = 0
		maxArrLen = 0
	)

	for i, v := range nums {
		if v == 0 {
			sum += -1
		} else {
			sum += 1
		}

		_, ok := sumToIndex[sum]
		if ok {
			maxArrLen = max(maxArrLen, i-sumToIndex[sum])
		} else {
			sumToIndex[sum] = i
		}
	}

	return maxArrLen
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
