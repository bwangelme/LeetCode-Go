package l80

/**
	复杂度分析：

		1. 时间复杂度: O(N)
		2. 空间复杂度: O(1) 额外申请的存储空间仅有两个变量
 */
func removeDuplicates(nums []int) int {
	idx := 0
	idxCount := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[idx] {
			idx++
			nums[idx] = nums[i]
			idxCount = 1
			continue
		}

		idxCount++

		// 相等且出现在两次以内
		if idxCount == 2 {
			idx++
			nums[idx] = nums[i]
		}
	}

	return idx + 1

}
