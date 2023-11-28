package l62

/*
## 题目

一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？

## 解题思路

动态规划的办法，转移方程

f(m, n) = f(m-1, n) + f(m, n-1)

到达 m,n 的路径是 到达 m-1,n 的路径+ 到达 m,n-1的路径

## 复杂度分析

时间复杂度和空间复杂度是 O(m * n)
*/

func uniquePaths(m int, n int) int {
	var dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 处于边上的时候，由于机器人只能向右或向下移动，达到边上方格的路径只有一条
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}
