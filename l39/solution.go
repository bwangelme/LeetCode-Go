package l39

/*
## 题目描述

给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。

candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。

对于给定的输入，保证和为 target 的不同组合数少于 150 个。

## 解题思路

利用回溯法来解题

遍历 candidates 中的每一个元素，每次可以将这个元素加入子集，或者不加入子集

由于子集中元素可以重复，因此加入后，下层循环索引不变，而可以继续加

递归的出口是，子集的和 > target, 当子集的和 == target 时，这是我们寻找的目标子集，加入结果集中

## 复杂度分析

N = len(candidates)
M = target 的值

- 空间复杂度

最坏的递归深度是 M, 此时 subset 的长度也是 M, 因此最坏空间复杂度是 O(M)

- 时间复杂度

整体遍历情况是一颗二叉树，二叉树的最坏深度是 M

最坏时间复杂度是 O(2^M)
*/

func combinationSum(candidates []int, target int) [][]int {
	var subset = make([]int, 0)
	var result = make([][]int, 0)
	helper(candidates, target, 0, subset, &result)
	return result
}

func helper(nums []int, target int, i int, subset []int, result *[][]int) {
	if target == 0 {
		resItem := make([]int, len(subset))
		copy(resItem, subset)
		*result = append(*result, resItem)
		return
	} else if target > 0 && i < len(nums) {
		// 不将当前数字加入到子集中
		helper(nums, target, i+1, subset, result)

		// 将当前数字加入到子集中
		item := nums[i]
		subset = append(subset, item)
		// 因为数字可以重复添加，因此将数字加入子集后，i 不变，下一层循环可以继续加
		helper(nums, target-nums[i], i, subset, result)
		subset = subset[0 : len(subset)-1]
	}
}
