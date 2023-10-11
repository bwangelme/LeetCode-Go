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
	firstGteRow, eq := searchFirstGteRow(matrix, target)
	if eq {
		return true
	}

	// 第一个比 target 大的行，不存在
	if firstGteRow == -1 {
		return false
	}

	targetRow := firstGteRow
	targetIdx := binarySearch(matrix[targetRow], target)
	return targetIdx != -1
}

// searchFirstGtRow
// 查找第一个 row[n] >= target 的行
//
// return: 行索引，targetRow[n-1] 是否等于 target
func searchFirstGteRow(matrix [][]int, target int) (int, bool) {
	var (
		m     = len(matrix)
		n     = len(matrix[0])
		left  = 0
		right = m - 1
	)

	for left <= right {
		mid := left + ((right - left) >> 1)
		midRowN := matrix[mid][n-1]
		if midRowN > target {
			if mid == 0 || matrix[mid-1][n-1] < target {
				return mid, false
			} else {
				right = mid - 1
			}
		} else if midRowN == target {
			return mid, true
		} else {
			left = mid + 1
		}
	}

	return -1, false
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
