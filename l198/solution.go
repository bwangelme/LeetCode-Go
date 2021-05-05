package l198

import "math"

// 动态规划解题
// Money(N) 第N间房屋的财富
// dp(N) 偷窃N间房屋所得的最大金额
// dp(N) = max(Money(N) + dp(N-2), dp(N-1))
// dp(1) = Money(1)
// dp(2) = max(Money(1), Money(2))

// 时间复杂度 O(N)
// 空间复杂度 O(N)
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var dp = make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = int(math.Max(float64(nums[0]), float64(nums[1])))

	for i := 2; i < len(nums); i++ {
		dp[i] = int(math.Max(float64(nums[i]+dp[i-2]), float64(dp[i-1])))
	}

	return dp[len(nums) - 1]
}
