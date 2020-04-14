package l343

import "math"

func integerBreak(n int) int {
	if n <= 2 {
		n = 2
	}

	var dp = make([]int, n+1)
	// dp[1] = 1 是一种特殊的初始情况
	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= n; i++ {
		for j := i - 1; j >= 1; j-- {
			num := math.Max(
				float64(dp[i]),
				math.Max(
					float64(j)*float64(dp[i-j]),
					float64(j)*float64(i-j),
				),
			)
			dp[i] = int(num)
		}
	}

	return dp[n]
}
