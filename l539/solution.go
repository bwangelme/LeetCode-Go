package l539

import (
	"strconv"
	"strings"
)

/*
## 题目

给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。


示例 1：

输入：timePoints = ["23:59","00:00"]
输出：1
示例 2：

输入：timePoints = ["00:00","23:59","00:00"]
输出：0


提示：

2 <= timePoints.length <= 2 * 10^4
timePoints[i] 格式为 "HH:MM"

## 解题思路

1. 将输入的时间映射到一个数组做的哈希表中 MinFlags
2. 遍历哈希表 MinFlags，寻找差的最小值 findHelper

## 复杂度分析

n == len(timePoints)

时间复杂度: O(N+1440) == O(N), 1440 是哈希表的长度
空间复杂度: O(1), len(MinFlags)

MinFlags 的长度是固定的 1440
*/

func findMinDifference(timePoints []string) int {
	if len(timePoints) <= 0 {
		return 0
	}
	var MinFlags [1440]bool

	for _, timePoint := range timePoints {
		copos := strings.Split(timePoint, ":")
		hour, _ := strconv.Atoi(copos[0])
		minute, _ := strconv.Atoi(copos[1])
		key := hour*60 + minute

		// 输入的时间有可能重复，如果重复了，直接返回 0
		exist := MinFlags[key]
		if exist {
			return 0
		}
		MinFlags[key] = true
	}

	return findHelper(MinFlags[:])
}

func findHelper(MinFlags []bool) int {
	var (
		first   = -1
		minDiff = len(MinFlags)
		last    = -1
		prev    = -1
	)

	for i := 0; i < len(MinFlags); i++ {
		if !MinFlags[i] {
			continue
		}

		if prev >= 0 {
			minDiff = min(minDiff, i-prev)
		}

		prev = i
		if first == -1 {
			first = i
		}
		last = i
	}

	return min(minDiff, first+len(MinFlags)-last)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
