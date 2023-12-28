package l15

import (
	"sort"
)

/*
## 复杂度分析

- N 表示 len(nums)
- 时间复杂度:
  - 排序: O(N * logN)
  - 针对每个 i 执行 twoSum: O(N^2)
  - 总复杂度: O(N*longN) + N(N^2)

- 空间复杂度:
  - 没有额外空间，result 可以认为是结果必须的，因此是 O(1)
*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var (
		n      = len(nums)
		result = make([][]int, 0)
	)
	for i := 0; i < n-2; {
		subRes := twoSum(nums, i)
		result = append(result, subRes...)

		temp := nums[i]
		for i < n && nums[i] == temp {
			i++
		}
	}

	return result
}

func twoSum(arr []int, i int) [][]int {
	var (
		j      = i + 1
		k      = len(arr) - 1
		result = make([][]int, 0)
	)

	for j < k {
		if arr[i]+arr[j]+arr[k] == 0 {
			result = append(result, []int{arr[i], arr[j], arr[k]})

			temp := arr[j]
			for j < k && arr[j] == temp {
				j++
			}
		} else if arr[i]+arr[j]+arr[k] < 0 {
			j++
		} else {
			k--
		}
	}

	return result
}
