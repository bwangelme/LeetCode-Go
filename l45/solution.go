package l45

/*
## 题目

给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。

每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:

    0 <= j <= nums[i]
    i + j < n

返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。


1 <= nums.length <= 10^4
0 <= nums[i] <= 1000
题目保证可以到达 nums[n-1]

## 解题思路

运用贪心的思路，每次求解的是当前局部最优解

用 i 从后往前遍历数组，每次都尽可能往前走最多的步数，直到走到0

针对每个 i, 用 j 从 0 循环到 i-1, j 在尽可能小的时候退出，这样就能每次让 i 尽可能走最多的步数

## 复杂度分析

N == len(nums)

- 时间复杂度 O(N)
- 空间复杂度 O(1)
*/

func jump(nums []int) int {
	var n = len(nums)
	if n <= 1 {
		return 0
	}

	var (
		count = 0
		i     = n - 1
	)

	for i > 0 {
		var (
			newI = -1
		)
		for j := 0; j < i; j++ {
			var delta = i - j
			maxLen := nums[j]
			if maxLen >= delta {
				newI = j
				break
			}
		}

		if newI == -1 {
			return -1
		} else {
			i = newI
		}
		count++
	}

	return count
}
