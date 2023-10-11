package l34

/**
## 题目
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

## 复杂度说明

BinarySearchFirst 和 BinarySearchLast 的复杂度都是 O(logN)

searchRange 的复杂度是 O(2 * logN) == O(logN)
**/

func searchRange(nums []int, target int) []int {
	first := BinarySearchFirst(nums, target)
	last := BinarySearchLast(nums, target)

	return []int{first, last}
}

func BinarySearchFirst(arr []int, num int) int {
	var (
		start = 0
		end   = len(arr) - 1
	)

	for start <= end {
		mid := start + ((end - start) >> 1)
		midNum := arr[mid]
		if midNum > num {
			end = mid - 1
			continue
		} else if midNum < num {
			start = mid + 1
			continue
		}

		if num == midNum {
			// start 是从0一直往上加的，所以 start 最小是0
			// 因为 start 最小是0, mid = start + (end-start)/2, end-start 最小是0, 所以 mid 最小是0, 不会出现 -1
			if mid == 0 || arr[mid-1] != num {
				return mid
			} else {
				end = mid - 1
			}
		}
	}

	return -1
}

func BinarySearchLast(arr []int, num int) int {
	var (
		start = 0
		n     = len(arr)
		end   = n - 1
	)

	for start <= end {
		mid := start + ((end - start) >> 1)
		midNum := arr[mid]
		if midNum > num {
			end = mid - 1
			continue
		} else if midNum < num {
			start = mid + 1
			continue
		}

		if num == midNum {
			if mid == n-1 || arr[mid+1] != num {
				return mid
			} else {
				start = mid + 1
			}
		}
	}

	return -1
}
