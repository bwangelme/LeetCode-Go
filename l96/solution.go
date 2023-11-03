package l96

/*
## 题目

给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。


    1 <= n <= 19

## 复杂度分析

n 表示输入的数量

时间复杂度是 O(n^2)
空间复杂度是 O(n)
*/

func numTrees(n int) int {
	G := make([]int, n+1)
	G[0] = 1
	G[1] = 1
	for N := 2; N <= n; N++ {
		for i := 1; i <= N; i++ {
			G[N] += G[i-1] * G[N-i]
		}
	}

	return G[n]
}
