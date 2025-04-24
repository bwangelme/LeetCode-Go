package l56

import (
	"sort"
)

/*
## 题目描述

以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

提示：

1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104

## 解题思路

1. 将 intervals 按照 start 排序
2. 遍历 intervals 中的每个集合，判断下一个集合的 start 是否 <= 当前集合的 end, 如果是，将两个集合合并，新集合的 end 为两个集合 end 的较大值

## 复杂度分析

N == len(intervals)

空间复杂度，新申请了 res 用来存储结果，最坏情况下空间复杂度为 O(N)

时间复杂度，排序的时间复杂度为 O(NlogN)

遍历 intervals 合并集合的时间复杂度为 O(N)

因此总的时间复杂度为 O(NlogN)
*/

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res = make([][]int, 0)

	var i = 0
	for i < len(intervals) {
		temp := []int{intervals[i][0], intervals[i][1]}
		var j = i + 1
		for ; j < len(intervals) && intervals[j][0] <= temp[1]; j++ {
			temp[1] = Max(temp[1], intervals[j][1])
		}

		res = append(res, temp)
		i = j
	}

	return res
}
