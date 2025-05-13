package l78

/*
## 题目描述

给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

## 解题思路

以一颗二叉树的形式来遍历整个集合

树的跟节点是空集合，每个节点有两个子节点，是否要添加集合中的一个数字到子集中，叶子节点表示一个结果

例如集合 [1, 2], 它对应的树就是这样的

[]
|---- 添加 1 [1]
		|---- 添加2 [1, 2]
		|---- 不添加2 [1]
|---- 不添加 1
		|---- 添加2 [2]
		|---- 不添加2 []

四个叶子节点就是四个子集

程序将叶子节点的所有结果添加到一个 result 中，然后将 result 返回

## 复杂度分析

N = len(nums)

- 空间复杂度，递归调用的深度是 N, 查到一个结果之后开始回溯
- 时间复杂度

这是一颗满二叉树，总节点的个数是 2^(N+1) -1 ，因此 helpers 函数会执行 2^(N+1) -1 次，因此时间复杂度是 O(2^N)

*/

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	var subset = make([]int, 0)
	var result = make([][]int, 0)
	helpers(nums, 0, subset, &result)

	return result
}

func helpers(nums []int, index int, subset []int, result *[][]int) {
	if index == len(nums) {
		resItem := make([]int, len(subset))
		copy(resItem, subset)
		*result = append(*result, resItem)
		return
	}

	helpers(nums, index+1, subset, result)

	item := nums[index]
	subset = append(subset, item)
	helpers(nums, index+1, subset, result)
	subset = subset[0 : len(subset)-1]
}
