package l122

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
func maxProfit(prices []int) int {
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
