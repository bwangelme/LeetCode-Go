package l121

import "math"

/*
## 题目

给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

## 题解

cost 表示买入价格，寻找最低的买入价格
prices[i] - cost 表示利润，因为卖出必须在买入之后，所以

prices[i] - cost 必须在算出了最低 cost 之后进行

求得的 max(prices[i] - cost) 就是最高的利润了

## 复杂度分析

N 表示 len(prices)

时间复杂度 O(N)
空间复杂度 O(1)

*/

func maxProfit(prices []int) int {
	n := len(prices)
	cost := math.MaxInt64
	profit := 0
	for i := 0; i < n; i++ {
		cost = min(cost, prices[i])
		profit = max(profit, prices[i]-cost)
	}

	return profit
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
