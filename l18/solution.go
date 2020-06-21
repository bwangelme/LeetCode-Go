package l18

import (
	"sort"
)

func rightSkipEqual(start, end int, nums []int) (startRes int) {
	for start < end {
		if nums[start] == nums[start+1] {
			start++
		} else {
			return start
		}
	}

	return start
}

func leftSkipEqual(start, end int, nums []int) (endRes int) {
	for start < end {
		if nums[end] == nums[end-1] {
			end--
		} else {
			return end
		}
	}

	return end
}

// 最坏时间复杂度分析
// 快排的时间复杂度是 O(n * logN)
// 选取结果的时间复杂度是 O((n-2) * (n-i-3) * (n-i-3)) ≈ O(n^3)
// 总的时间复杂度是 O(n * logN) + O(n^3)

// 运行结果
// Runtime  Memory
// 32 ms	2.8 MB
func fourSum(nums []int, target int) (result [][]int) {
	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}

			start := j + 1
			end := len(nums) - 1
			expectedSum := target - nums[i] - nums[j]

			for start < end {
				sum := nums[start] + nums[end]
				if sum == expectedSum {
					item := []int{nums[i], nums[j], nums[start], nums[end]}
					result = append(result, item)
					start = rightSkipEqual(start, end, nums)
					start++
					end = leftSkipEqual(start, end, nums)
					end--
				} else if sum < expectedSum {
					start = rightSkipEqual(start, end, nums)
					start++
				} else if sum > expectedSum {
					end = leftSkipEqual(start, end, nums)
					end--
				}
			}
		}
	}

	return result
}
