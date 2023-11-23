package l740

import (
	"math"
)

// 将 nums 转换成 trans 数组后，这个问题就变成了 198打家劫舍 问题
// 选择了 trans 中的一个点后，它的前后都不能选择了，这样就可以推导出动态规划方程
// dp[i] 表示在 [0, i] 之间能拿到的最大点数
// dp[N] = max(trans[N] + dp[N-2], dp[N-1])
// dp[0] = trans[0]
// dp[1] = trans[1] // 因为 trans[0] = 0，所以 dp[1] = max(trans[0], trans[1]) = trans[1]
func deleteAndEarn(nums []int) int {
	var trans = make(map[int]int)
	for _, num := range nums {
		trans[num] += num
	}

	var dp = make([]int, 10001)

	dp[0] = 0
	dp[1] = trans[1]
	for i := 2; i < len(dp); i++ {
		dp[i] = int(math.Max(float64(dp[i-1]), float64(dp[i-2]+trans[i])))
	}

	return dp[len(dp)-1]
}
