package l33

/**
## 题目

整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

给你旋转后的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。

提示：

- 1 <= nums.length <= 5000
- -104 <= nums[i] <= 104
- nums 中的每个值都 独一无二
- 题目数据保证 nums 在预先未知的某个下标上进行了旋转
- -104 <= target <= 104

## 复杂度分析

findMinIdx 的复杂度是 O(logN)

两次 binSearch 总共查找的元素是 N, 所以 两次 binSearch 的复杂度是 O(logN)

总复杂度是 O(2 * logN)，即是 O(logN)

*/

func search(nums []int, target int) int {
	minIdx := findMinIdx(nums)
	targetIdx := binSearch(nums, 0, minIdx-1, target)
	if targetIdx != -1 {
		return targetIdx
	}
	targetIdx = binSearch(nums, minIdx, len(nums)-1, target)
	return targetIdx
}

func binSearch(nums []int, left, right int, target int) int {
	var (
		mid int
	)

	for left <= right {
		mid = left + ((right - left) >> 1)
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// findMinIdx
// 寻找数组中最小值的索引
func findMinIdx(nums []int) int {
	var (
		n     = len(nums)
		left  = 0
		right = n - 1
		mid   int
	)

	for left <= right {
		mid = left + ((right - left) >> 1)
		/**
		 为了跳出循环，需要让 right < left, 最后一次 left,right 指向同一个最小的数，改变 right 满足循环退出条件, 所以这里是个 <= 的条件

		1. min == nums[n-1], 此时改变 left 或 right 都可以
		2. min < nums[n-1], 此时最后一次只能改变 right

		所以最终我们需要改变 right, 让循环跳出，left 保留在最小值上
		**/
		if nums[mid] <= nums[n-1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
