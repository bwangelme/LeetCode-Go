package l70

// 时间复杂度 O(N)
// 空间复杂度 O(N)
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	var dp = make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
