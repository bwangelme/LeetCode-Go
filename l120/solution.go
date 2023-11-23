package l120

/*
## 题目

给定一个三角形 triangle ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。


1 <= triangle.length <= 200
triangle[0].length == 1
triangle[i].length == triangle[i - 1].length + 1
-10^4 <= triangle[i][j] <= 10^4

## 解题思路

这是典型的动态规划问题，可以分为四步解题

1. 定义状态

minSum[i][j] 表示从三角形最底层到 i, j ，这个路径中走的路径的最小和

2. 状态转移方程

minSum[i][j] = min(minSum[i+1][j], minSum[i+1][j+1]) + triangle[i][j]

根据最小值决定是从 j 还是 j+1 走上来

3. 初始状态

最底层每个数字的最小路径和都是自身

minSum[n-1][j] = triangle[n-1][j]

4. 最终求值

求最终走到顶部节点的最小路径和 minSum[0][0]

## 复杂度分析

N 表示三角形的高度，及 len(triangle)

### 空间复杂度

申请了额外的数组，minSum, 因此是 O(N^2)

### 时间复杂度

// 初始化 minSum
O(N)

// 求动态规划中的初始条件
O(N)

// 求动态规划方程的状态
O(N^2)

因此时间复杂度是 O(2N) + O(N^2)

*/

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	var MAX = n

	// 从底向上遍历
	minSum := make([][]int, MAX)
	for i := 0; i < len(minSum); i++ {
		minSum[i] = make([]int, MAX)
	}

	// 初始条件
	for i := 0; i < len(triangle[n-1]); i++ {
		minSum[n-1][i] = triangle[n-1][i]
	}

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			minSum[i][j] = min(minSum[i+1][j], minSum[i+1][j+1]) + triangle[i][j]
		}
	}

	return minSum[0][0]
}

// minimumTotalV2
/*
## 解题思路

这个程序的思路和 minimumTotalV2 类似，也是自底向上算每个数的最小路径

但是这个直接用 triangle, 没有新申请空间，空间复杂度是 O(1), 时间复杂度是 O(n^2)

https://passage-1253400711.cos.ap-beijing.myqcloud.com/2023-11-23-100900.png
*/
func minimumTotalV2(triangle [][]int) int {
	n := len(triangle)
	for i := n - 1; i >= 0; i-- {
		for j := 1; j < len(triangle[i]); j++ {
			triangle[i-1][j-1] = triangle[i-1][j-1] + min(triangle[i][j], triangle[i][j-1])
		}
	}

	return triangle[0][0]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
