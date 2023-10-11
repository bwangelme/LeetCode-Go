package l74

/**
## 题目

给你一个满足下述两条属性的 m x n 整数矩阵：

每行中的整数从左到右按非递减顺序排列。
每行的第一个整数大于前一行的最后一个整数。
给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。

## 复杂度分析

searchFirstGtRow 查找 target 的下一行，复杂度是 O(logM)
binarySearch 在某一行中查找 target，复杂度是 O(logN)

时间复杂度是 O(log(M + N))
**/

func searchMatrix(matrix [][]int, target int) bool {
	firstGtRow := searchFirstGtRow(matrix, target)
	// 第一个比 target 大的行，是首行或者不存在
	if firstGtRow == -1 || firstGtRow == 0 {
		return false
	}

	targetRow := firstGtRow - 1
	targetIdx := binarySearch(matrix[targetRow], target)
	return targetIdx != -1
}

// searchFirstGtRow
// 查找第一个 row[0] > target 的行
func searchFirstGtRow(matrix [][]int, target int) int {
	var (
		m     = len(matrix)
		left  = 0
		right = m - 1
	)

	for left <= right {
		mid := left + ((right - left) >> 1)
		midRowFirst := matrix[mid][0]
		if midRowFirst > target {
			if mid == 0 || matrix[mid-1][0] <= target {
				return mid
			} else {
				right = mid - 1
			}
		} else {
			left = mid + 1
		}
	}

	return -1
}

// binarySearch
// 利用二分查找法在数组中查找 target
func binarySearch(nums []int, target int) int {
	var (
		n     = len(nums)
		left  = 0
		right = n - 1
	)
	for left <= right {
		mid := left + ((right - left) >> 1)
		midNum := nums[mid]
		if midNum > target {
			right = mid - 1
		} else if midNum < target {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
