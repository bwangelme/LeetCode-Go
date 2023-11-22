package l122

/*
## 题目

给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。

在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。

返回 你能获得的 最大 利润 。

1 <= prices.length <= 3 * 104
0 <= prices[i] <= 104

## 解题思路

利用动态规划的思路

1. 定义状态
	dp[i][0] 第 i 天持有现金的最大收益
	dp[i][1] 第 i 天持有股票的最大收益
2. 定义状态转移方程
	dp[i][0] = max(dp[i-1][1] + p[i], dp[i-1][0])
	dp[i][1] = max(dp[i-1][1], dp[i-1][0]-p[i])

	这个状态转移方程就是 i 天确定 i-1 天的操作，今天想要持有现金或股票，决定昨天是买股票，卖出股票，还是等待
3. 定义初始状态
	dp[0][0] = 0
	dp[0][1] = -p[0]
4. 输出结果
	最后输出 dp[n-1][0], 最后一天持有现金的收益
*/

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	var dp = make([][2]int, n)

	// 3
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	// 2
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	}

	// 4
	return dp[n-1][0]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

/*
尝试用贪心算法来解题，每天往后找收益最大的一天，局部最优解看能否成为全局最优解

试了一下，发现不行

7 1 5 3 6 4

这个解法算出来的解是

buy 1
sell 5
buy 5
sell 6

得出来的最高利润是 5
*/
func maxProfitErr(prices []int) int {
	n := len(prices)
	var sum = 0
	for i := 0; i < n; {
		buy := prices[i]
		buyDay := i
		var mProfitPerDay float64
		var mProfit int
		for j := i + 1; j < n; j++ {
			profit := prices[j] - buy
			deltaDay := j - buyDay
			profitPerDay := float64(profit) / float64(deltaDay)
			if profitPerDay > mProfitPerDay {
				mProfitPerDay = profitPerDay
				mProfit = profit
				i = j
			}
		}
		if mProfitPerDay == 0 {
			i++
		} else {
			sum += mProfit
		}
	}

	return sum
}
