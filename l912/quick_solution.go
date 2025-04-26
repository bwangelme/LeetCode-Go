package l912

import "math/rand"

/*
这是不超时的答案
*/
func sortArrayFast(nums []int) []int {
	quickSortF(nums)
	return nums
}

func quickSortF(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	// 如果你想随机选择一个索引作为基准值
	pivotIndex := rand.Intn(n)
	nums[0], nums[pivotIndex] = nums[pivotIndex], nums[0]

	pivot := nums[0]
	i, j := -1, n

	for i < j {
		for i++; nums[i] < pivot; i++ {
		}
		for j--; nums[j] > pivot; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	quickSortF(nums[:j+1])
	quickSortF(nums[j+1:])
}
