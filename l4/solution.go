package l4

import (
	"math"
)

/*
	题目:
	给定两个排序的数组，nums1 和 nums2，它们的大小分别是 m 和 n，返回这两个数组的中位数
	注意: 最坏时间复杂度应该是 O(log(m + n))

	复杂度分析:
	1. 时间复杂度是 O(log(m+n))，每次递归问题规模都变小了
	2. 空间复杂度是 O(log(m+n)) 函数被调用了 log(m+n) 次
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	total := m + n

	if total%2 == 1 {
		return float64(findKth(nums1, m, nums2, n, total/2+1))
	} else {
		m1 := findKth(nums1, m, nums2, n, total/2)
		m2 := findKth(nums1, m, nums2, n, total/2+1)
		return float64(m1+m2) / 2.0
	}
}

func findKth(numsA []int, m int, numsB []int, n int, k int) int {
	if m > n {
		return findKth(numsB, n, numsA, m, k)
	}
	if m == 0 {
		return numsB[k-1]
	}
	if k == 1 {
		return int(math.Min(float64(numsA[0]), float64(numsB[0])))
	}

	ia := (k + 1) / 2
	if ia > m {
		ia = m
	}
	ib := k - ia
	if numsA[ia-1] < numsB[ib-1] {
		return findKth(numsA[ia:], m-ia, numsB, n, k-ia)
	} else if numsA[ia-1] > numsB[ib-1] {
		return findKth(numsA, m, numsB[ib:], n-ib, k-ib)
	} else {
		return numsA[ia-1]
	}
}
