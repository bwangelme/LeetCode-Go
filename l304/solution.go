package l304

/*
## 题目

给定一个二维矩阵 matrix，以下类型的多个请求：

计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1, col1) ，右下角 为 (row2, col2) 。
实现 NumMatrix 类：

NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
int sumRegion(int row1, int col1, int row2, int col2) 返回 左上角 (row1, col1) 、右下角 (row2, col2) 所描述的子矩阵的元素 总和 。

提示：

m == matrix.length
n == matrix[i].length
1 <= m, n <= 200
-10^5 <= matrix[i][j] <= 10^5
0 <= row1 <= row2 < m
0 <= col1 <= col2 < n
最多调用 10^4 次 sumRegion 方法

## 解题思路

输入矩阵 matrix[m][n]，求任一子矩阵的和

定义一个辅助矩阵 sum[m+1][n+1], 令

sum[i+1][j+1] = Sum(matrix[0][0], ....  matrix[i][j])

sum 比 matrix 多一行一列，是为了计算首行首列时方便

任一子矩阵 matrix[r1][c1] .. matrix[r2][c2] 的和，可以变成四个以 (0,0) 为左上角的矩阵的运算

sum[r2+1][c2+1] - sum[r1][c2+1] - sum[r2+1][c1] + sum[c1][r1]

sum 辅助矩阵的计算方法是

for i := range(0, m-1) {
	rowSum = 0
	for j := range(0, n-1) {
		rowSum += matrix[i][j]
		sum[i+1][j+1] = sum[i][j+1] + rowSum
	}
}

## 复杂度分析

m 表示矩阵的行数
n 表示矩阵的列数

时间复杂度:

- Constructor: O(m * n)
- SumRegion: O(1)

空间复杂度:

创建了额外的 sum 数组, O(m * n)
*/

type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}

	m := len(matrix)
	n := len(matrix[0])

	sum := make([][]int, m+1)
	for j, _ := range sum {
		sum[j] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		rowSum := 0
		for j := 0; j < n; j++ {
			rowSum += matrix[i][j]
			sum[i+1][j+1] = sum[i][j+1] + rowSum
		}
	}

	return NumMatrix{
		sum: sum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	val := this.sum[row2+1][col2+1] - this.sum[row1][col2+1] - this.sum[row2+1][col1] + this.sum[row1][col1]
	return val
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
