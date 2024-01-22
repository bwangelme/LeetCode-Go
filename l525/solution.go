package l525

func findMaxLength(nums []int) int {
	// S[i] 表示 nums[0,i] 所有数字的和
	// sumToIndex 的键是 S[i], 值是第一次遇到 S[i] 时 i 的索引
	sumToIndex := make(map[int]int)

	// 如果 S[i] == 0, 那么表示 nums[0,i] 就是符合条件的数组, len(nums[0,i]) == i+1 == i - (-1)
	// 因此 sumToIndex[0] = -1
	sumToIndex[0] = -1

	var (
		sum       = 0
		maxArrLen = 0
	)

	for i, v := range nums {
		if v == 0 {
			sum += -1
		} else {
			sum += 1
		}

		_, ok := sumToIndex[sum]
		if ok {
			maxArrLen = max(maxArrLen, i-sumToIndex[sum])
		} else {
			sumToIndex[sum] = i
		}
	}

	return maxArrLen
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
