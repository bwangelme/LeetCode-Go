package l169

/*
## 题目描述
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1：

输入：nums = [3,2,3]
输出：3

示例 2：

输入：nums = [2,2,1,1,1,2,2]
输出：2

## 解题思路

遍历数组，用一个变量 candidate 记录当前的候选元素，用一个变量 count 记录当前的候选元素的出现次数。

当 count == 0 的时候，让 candidate 等于当前的元素，count 等于 1，
相当于让当前元素参选，后续元素进行投票(相同+1，不同-1)，如果当前元素被选下去了(count=0)，后续元素重新参选

## 复杂度分析

N = len(nums)

时间复杂度：O(N)
空间复杂度：O(1)
*/
func majorityElement(nums []int) int {
	var (
		candidate int
		count     int
	)
	for _, num := range nums {
		if count == 0 {
			candidate = num
			count = 1
			continue
		}

		if num == candidate {
			candidate = num
			count += 1
		} else {
			count -= 1
		}
	}

	return candidate
}
