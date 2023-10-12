package l69

/**
## 题目
给你一个非负整数 x ，计算并返回 x 的 算术平方根 。

由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。

## 复杂度分析

这个题目的解题思路还是二分查找，因为是计算平方，所以当知道 x > 2 的时候，我们可以只判断 0-x/2 的数字

时间复杂度是 O(logN/2) = O(logN)
*/

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	if x <= 2 {
		return 1
	}

	var (
		left  = 1
		right = x / 2
	)

	for left <= right {
		mid := left + ((right - left) >> 1)
		if mid*mid > x {
			right = mid - 1
		} else {
			// 因为相等的时候我移动了 left，所以我在函数尾部 return right
			left = mid + 1
		}
	}

	return right
}
