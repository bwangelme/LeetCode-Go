package l704

/*
## 题目说明
	给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
	链接：https://leetcode.cn/problems/binary-search
*/
func search(nums []int, target int) int {
	// 定义 target 在左闭右闭的区间 [start, end]
	start, end := 0, len(nums)-1
	for {
		// start == end 时, target 仍然在区间中, 不跳出循环
		if start > end {
			break
		}

		// 防止溢出
		middle := start + (end-start)/2
		if target > nums[middle] {
			start = middle + 1
		} else if target < nums[middle] {
			end = middle - 1
		} else {
			return middle
		}
	}

	return -1
}
