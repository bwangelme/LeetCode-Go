package l875

import (
	"math"
)

/*
## 题目描述

珂珂喜欢吃香蕉。这里有 n 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 h 小时后回来。

珂珂可以决定她吃香蕉的速度 k （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 k 根。如果这堆香蕉少于 k 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。

珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。

返回她可以在 h 小时内吃掉所有香蕉的最小速度 k（k 为整数）。

## 解题思路

left = 1
right = piles 中的最大值

在这个范围内进行二分查找，查找到 mid 后的到速度，根据速度计算吃完香蕉需要的时间h,
	如果计算的时间 h 小于目标时间 H,那么将速度调慢，让 h 变大一些，此时 right = mid -1
	如果计算的时间 h 大于目标时间 H, 那么将速度调快，让 h 便小一些，此时 left = mid+1

## 复杂度分析

N = len(piles)
M = max(piles)

二分查找的时间复杂度是 O(logM)
每次比较时, 计算以当前速度吃完花费时间H, 复杂度是 O(N)

因此总的时间复杂度是 O(N * logM)
*/

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minEatingSpeed(piles []int, h int) int {
	var max = math.MinInt
	for _, p := range piles {
		max = Max(p, max)
	}

	var (
		left  = 1
		right = max
	)

	for left <= right {
		mid := left + (right-left)/2
		hours := getHours(piles, mid)
		if hours <= h {
			if mid == 1 || getHours(piles, mid-1) > h {
				return mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func getHours(piles []int, speed int) int {
	var hours = 0
	for _, p := range piles {
		hours += (p + speed - 1) / speed
	}

	return hours
}
