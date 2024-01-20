package l713

/*
## 题目描述

给你一个整数数组 nums 和一个整数 k ，请你返回子数组内所有元素的乘积严格小于 k 的连续子数组的数目。


示例 1：

输入：nums = [10,5,2,6], k = 100
输出：8
解释：8 个乘积小于 100 的子数组分别为：[10]、[5]、[2],、[6]、[10,5]、[5,2]、[2,6]、[5,2,6]。
需要注意的是 [10,5,2] 并不是乘积小于 100 的子数组。
示例 2：

输入：nums = [1,2,3], k = 0
输出：0

提示:

1 <= nums.length <= 3 * 10^4
1 <= nums[i] <= 1000
0 <= k <= 10^6

## 解题思路

和 l209 的思路类似，因为子数组是连续的，就用 l, r 从 0 到 n-1 遍历数组，符合条件的数组 count += 1


关于以下代码的解释

		if l <= r {
			count += r + 1 - l
		}

令 f(l, r) 表示 nums[l:r] 中乘积小于 k 的子数组的个数

则

f(l, r+1) = f(l, r) + r+1-l

因为 nums[l:r] 之间所有的子数组已经统计过了，要统计新增的符合条件的子数组，都是以 r+1 为最后一个数

共有这些数组

nums[l:r+1], nums[l+1:r+1], .... nums[r:r+1], nums[r+1:r+1]

这一共是 r+1 - l 个数组

因此 count += r + 1 - l
*/

func numSubarrayProductLessThanK(nums []int, k int) int {
	var (
		l       = 0
		r       = 0
		product = 1
		count   = 0
	)
	for r = 0; r < len(nums); r++ {
		product *= nums[r]
		for l <= r && product >= k {
			product /= nums[l]
			l++
		}
		if l <= r {
			count += r + 1 - l
		}
	}

	return count
}
