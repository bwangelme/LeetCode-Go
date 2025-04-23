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

/*
## 题目描述

给你一个非负整数 x ，计算并返回 x 的 算术平方根 。

由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。

示例 1：

输入：x = 4
输出：2
示例 2：

输入：x = 8
输出：2
解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。

提示：

0 <= x <= 231 - 1

## 解题思路

从 1-x 进行二分查找，查找 target^2 > n && (target+1) ^ 2 < n

如果输入的 x 是0, 那么 left > right, 直接返回0就好

## 复杂度分析

N = x 的大小

空间复杂度 O(1)
时间复杂度 O(logN)
*/
func mySqrtV2(x int) int {
	var (
		left  = 1
		right = x
	)
	for left <= right {
		var mid = (left + right) / 2
		if mid <= x/mid {
			if (mid + 1) > x/(mid+1) {
				return mid
			}

			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return 0
}
