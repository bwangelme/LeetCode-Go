package sort

// TopK
/**
## 题目

O(n) 时间复杂度内求数组的第 k 大元素

## 思路介绍

这个题目也使用了分治的思想

利用快排的 partition 函数，获取最后一个数在数组中的大小索引 p(注意 p 是从0开始的)，并按大小左右分组

如果 p+1 > k, 则说明第 k 个数在左边, 在 start:p-1 中继续寻找
如果 p+1 < k, 则说明第 k 个数在右边，在 p+1:end 中继续寻找

如果 p+1 == k, 则是找到了目标

## 时间复杂度分析

T(1) = C
topKC 函数每次只查一边，另外一边不查，所以是 or
T(n) = T(sub_ln) or T(sub_rn)

假设 p 每次正好在中间点上，这样遍历完之后，查找的次数是

n + n/2 + n/4 + n/8 + ... 1

这是个等比数列，S_n = 2*n - 1, 用大 O 表示法是 O(n)

## 另一种解题思路

每次找到最小值，放到数组最前面，执行 k 次，那么第 k 个数就是第 k 大的数

这种思路中，时间复杂度是 O(k * n)

## 空间复杂度

本题中没有根据数组长度 N 额外申请空间，所以空间复杂度是 O(C)

如果考虑到函数栈的空间，那么一共调用了 O(logN) 次函数，空间复杂度是 O(logN)

*/
func TopK(arr []int64, k int) int64 {
	n := len(arr)
	if k > n {
		return -1
	}

	return topKC(arr, 0, n-1, k)
}

func topKC(arr []int64, start, end int, k int) int64 {
	p := partition(arr, start, end)

	if p+1 == k {
		return arr[p]
	}

	if p+1 > k {
		return topKC(arr, start, p-1, k)
	} else {
		return topKC(arr, p+1, end, k)
	}
}
