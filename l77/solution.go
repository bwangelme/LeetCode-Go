package l77

/*
## 题目描述

给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

你可以按 任何顺序 返回答案。

## 解题思路

以一颗二叉树的形式来遍历 1-n 个证书

树的跟节点是空集合，每个节点有两个子节点，是否要添加集合中的一个数字到子集中, 当子集中的元素数量达到 k 个时，就是一个目标结果

例如整数 3, 它对应的树就是这样的

[]
|---- 添加 1 [1]
		|---- 添加2 [1, 2]
				|---- 添加3 [1, 2, 3]
				|---- 不添加3 [1, 2]
		|---- 不添加2 [1]
				|---- 添加3 [1, 3]
				|---- 不添加3 [1]
|---- 不添加 1
		|---- 添加2 [2]
				|---- 添加3 [2, 3]
				|---- 不添加3 [2]
		|---- 不添加2 []
				|---- 添加3 [3]
				|---- 不添加3 [2]

我们期望将数量=2的子集求出来, 可以进行的剪支是，当当前集合数量等于2时，不需要向下递归了

## 复杂度分析

N = 参数 n 的大小

- 空间复杂度

递归的深度是 N, 因此空间复杂度是 O(N)

- 时间复杂度

遍历了满二叉树，在某些节点进行了剪支，但总体复杂度还是 O(2^N)

*/

func combine(n int, k int) [][]int {
	if k > n {
		return nil
	}

	var subset = make([]int, 0)
	var result = make([][]int, 0)
	helpers(n, k, 1, subset, &result)

	return result
}

func helpers(n int, k int, index int, subset []int, result *[][]int) {
	if k == len(subset) {
		resItem := make([]int, len(subset))
		copy(resItem, subset)
		*result = append(*result, resItem)
		return
	} else if index <= n {
		helpers(n, k, index+1, subset, result)

		subset = append(subset, index)
		helpers(n, k, index+1, subset, result)
		subset = subset[0 : len(subset)-1]
	}
}
