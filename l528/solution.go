package l528

import (
	"math/rand"
	"time"
)

/*
## 题目描述

给你一个 下标从 0 开始 的正整数数组 w ，其中 w[i] 代表第 i 个下标的权重。

请你实现一个函数 pickIndex ，它可以 随机地 从范围 [0, w.length - 1] 内（含 0 和 w.length - 1）选出并返回一个下标。选取下标 i 的 概率 为 w[i] / sum(w) 。

例如，对于 w = [1, 3]，挑选下标 0 的概率为 1 / (1 + 3) = 0.25 （即，25%），而选取下标 1 的概率为 3 / (1 + 3) = 0.75（即，75%）。


示例 1：

输入：
["Solution","pickIndex"]
[[[1]],[]]
输出：
[null,0]
解释：
Solution solution = new Solution([1]);
solution.pickIndex(); // 返回 0，因为数组中只有一个元素，所以唯一的选择是返回下标 0。
示例 2：

输入：
["Solution","pickIndex","pickIndex","pickIndex","pickIndex","pickIndex"]
[[[1,3]],[],[],[],[],[]]
输出：
[null,1,1,1,1,0]
解释：
Solution solution = new Solution([1, 3]);
solution.pickIndex(); // 返回 1，返回下标 1，返回该下标概率为 3/4 。
solution.pickIndex(); // 返回 1
solution.pickIndex(); // 返回 1
solution.pickIndex(); // 返回 1
solution.pickIndex(); // 返回 0，返回下标 0，返回该下标概率为 1/4 。

由于这是一个随机问题，允许多个答案，因此下列输出都可以被认为是正确的:
[null,1,1,1,1,0]
[null,1,1,1,1,1]
[null,1,1,1,0,0]
[null,1,1,1,0,1]
[null,1,0,1,0,0]
......
诸若此类。


提示：

1 <= w.length <= 104
1 <= w[i] <= 105
pickIndex 将被调用不超过 104 次

## 解题思路

构造一个 sums 数组，存储每个下标的权重之和

例如 w 是 1 2 3 4
那么 sums 就是 1 3 6 10

然后在 [1, 10] 之间选择随机数, 判断随机数落的区间

如果是 (1, 1], 那么就选择下标0
如果是 (1, 3], 那么选择下标1
如果是 (3, 6], 那么选择下标2
以此类推

寻找随机数落入的区间，需要使用二分法，
查找第一个大于随机数的索引，如果他的前一个也小于，那么就是这个数就是要找的区间的右半部分，就是要返回的权重

## 复杂度分析

N = len(w)

空间复杂度，使用了额外的数组 sums,因此空间复杂度是 O(N)

时间复杂度

Constructor 函数的时间复杂度是 O(N)
PickIndex 函数的时间复杂度是 O(logN)


*/

type Solution struct {
	total int
	sums  []int
}

func Constructor(w []int) Solution {
	sums := make([]int, len(w))
	var total = 0
	for i := 0; i < len(w); i++ {
		total += w[i]
		sums[i] = total
	}

	return Solution{
		total: total,
		sums:  sums,
	}
}

func (this *Solution) PickIndex() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := r.Intn(this.total) + 1

	var (
		low  = 0
		high = len(this.sums) - 1
	)
	for low <= high {
		mid := (low + high) / 2
		// 寻找第一个比 p 大的数字
		if this.sums[mid] >= p {
			if mid == 0 || this.sums[mid-1] < p {
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

/**
 * Your Solution object will be instantiated and called as such:
 */
