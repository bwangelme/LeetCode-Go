package l26

/*
	题目要求: 不申请任何数组空间，申请的空间应该是O(1)

	复杂度分析:
		1. 空间复杂度是 O(1)
		2. 时间复杂度是 O(N)
*/
func removeDuplicates(nums []int) int {
	idx := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[idx] {
			idx++
			nums[idx] = nums[i]
		}
	}

	return idx + 1
}
