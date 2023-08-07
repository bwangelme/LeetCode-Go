package l9

/*
## 题目

给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
例如，121 是回文，而 123 不是。
*/

import (
	"strconv"
)

/*
## 复杂度分析

时间复杂度: O(N) N 是 X 的大小
空间复杂度: O(N) N 和字符串占用的空间比为 N/10, 故复杂度为 O(N)
*/
func isPalindrome(x int) bool {
	xStr := strconv.FormatInt(int64(x), 10)

	for idx := 0; idx < len(xStr)/2; idx++ {
		revertIdx := len(xStr) - 1 - idx
		if xStr[idx] != xStr[revertIdx] {
			return false
		}
	}

	return true
}

/*
## 复杂度分析

时间复杂度: O(N/10 * 2) == O(N)
空间复杂度: O(N/10) == O(N)
*/
func isPalindromeNoStr(x int) bool {
	if x < 0 {
		return false
	}

	var nums = make([]int, 0)

	for x > 0 {
		n := x % 10
		nums = append(nums, n)
		x = x / 10
	}

	for idx := 0; idx < len(nums)/2; idx++ {
		revertIdx := len(nums) - 1 - idx

		if nums[idx] != nums[revertIdx] {
			return false
		}
	}

	return true
}
