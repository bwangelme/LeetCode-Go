package l70

/*
## 题目

假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

## 解题思路

为了到达 N 级台阶，最后一步只有两种办法

1. 还剩一级台阶, 总方法数 f(n) = f(n-1)
2. 还剩两极台阶，总方法数 f(n) = f(n-2)

所以到达 N 级台阶的方法就是

f(n) = f(n-1) + f(n-2)

这和斐波那契数列类似，只不过初始值不同

跳台阶的初始值

f(1) = 1
f(2) = 2

求 N 级台阶的跳法即是求 f(n)
*/

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

// 时间复杂度 O(N)
// 空间复杂度 O(1)
func climbStairsV1(n int) int {
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	var (
		a   = 1
		b   = 2
		sum = b + a
	)
	for i := 3; i <= n; i++ {
		sum = b + a
		a = b
		b = sum
	}

	return sum
}
