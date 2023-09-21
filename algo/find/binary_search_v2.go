package find

/**
二分查找时，low 是从小到大移动， high 是从大到小移动。

根据 low 和 high 移动的方向，我们可以提出两个二分查找的变种问题

1. 查找第一个大于某数字的值的索引
2. 查找最后一个小于某数字的值的索引

而这两个问题的反面问题

1. 查找最后一个大于某数字的值的索引
2. 查找第一个小于某数字的值的索引

这两个问题都不用二分查找，直接比较最后一个值/第一个值就好, 答案存在就是第一个/最后一个
*/

//BinarySearchFirstGt
/**
## 题目

查找数组中第一个 大于 指定值的数的索引

## 解题思路

从 high 向 low 压缩，每次 high 压缩的时候，判断 high 是不是这个数

## 复杂度分析

最坏时间复杂度是 O(logN), N 为数组长度
*/
func BinarySearchFirstGt(arr []int64, num int64) int {
	var (
		n    = len(arr)
		low  = 0
		high = n - 1
	)

	for low <= high {
		mid := low + ((high - low) >> 1)
		midNum := arr[mid]
		if midNum > num {
			if mid == 0 || arr[mid-1] <= num {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}

	return -1
}

//BinarySearchLastLt
/**
## 题目

查找最后一个小于指定数的数的索引

## 思路

low 不断网上压缩，low 压缩时判断 low 是否是最后一个小于指定数

## 复杂度分析

时间复杂度是 O(logN)
*/
func BinarySearchLastLt(arr []int64, num int64) int {
	var (
		n    = len(arr)
		low  = 0
		high = n - 1
	)

	for low <= high {
		mid := low + ((high - low) >> 1)
		midNum := arr[mid]
		if midNum < num {
			if mid == n-1 || arr[mid+1] >= num {
				return mid
			} else {
				low = mid + 1
			}
		} else {
			high = mid - 1
		}
	}

	return -1
}
