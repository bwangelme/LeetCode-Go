package l912

import "math/rand"

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	if end > start {
		pivot := partition(nums, start, end)
		quickSort(nums, start, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}

func partition(nums []int, start, end int) int {
	p := rand.Intn(end-start+1) + start
	swap(nums, p, end)
	var small = start - 1

	for i := start; i < end; i++ {
		if nums[i] < nums[end] {
			small++
			swap(nums, small, i)
		}
	}

	small++
	swap(nums, small, end)

	return small
}

func swap(nums []int, idx1, idx2 int) {
	if idx1 != idx2 {
		temp := nums[idx1]
		nums[idx1] = nums[idx2]
		nums[idx2] = temp
	}
}
