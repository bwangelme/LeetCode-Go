package l56

import (
	"sort"
)

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
