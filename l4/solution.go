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

/**
  说明: 寻找合并后且排好序的数组 numsC(numsC = sorted(numsA + numsB)) 中排在第 k 位的元素

  实现原理:

  ## 递推公式

    从 numsA 中选取前 i 个元素(i < k)，从 numsB 中选取前 (k - i) 个元素
  1. 如果 numsA[i-1] > numsB[k-i-1]，那么 numsC 的前 k 个元素中一定有 numsB[0:k-i-1]
  2. 同理，如果 numsA[i-1] < numsB[k-i-1]，那么 numsC 的前 k 个元素中一定有 numsA[0:i-1]
  3. 如果 numsA[i-1] == numsB[k-i-1]，那么 len(numsA[0:i-1] + numsB[0:k-i-1]) == k，则
  numsA[i-1] 是 numsC 中的第 k 个元素。

    那么就可以将问题规模缩小，变成

    1. 当 numsA[i-1] > numsB[k-i-1]时，
       寻找 numsB[k-i:len(numsB)] + numsA 中排在第 i 位的元素
    2. 当 numsA[i-1] < numsB[k-i-1]时,
       寻找 numsA[i:len(numsA)] + numsB 中排在第 k - i 位的元素

    上述递归过程可以简述为 __谁小切谁__ 。

  ## 递归出口

    0. 因为 len(numsA) 始终小于 len(numsB)，所以当 len(numsA) == 0 的时候，
       排在第k位的数就是 numsB[k-1]
    1. 当 k == 1 的时候，排在第一位的数就是 min(numsA[0], numsB[0])
    2. 递推公式中论述的 numsA[i-1] == numsB[k-i-1] 也是一个递归出口

  ## 算法复杂度分析

    这个程序使用递归来解决，并且每次递归都会使问题规模都会变为原来的1/N，所以最坏时间复杂度是 O(log(M + N))

    这个递归程序的时间复杂度分析可以用二叉查找树来考虑，
    在 [0, N) 之间二分查找一个一个数字，叶子节点共有 N 个，则树的高度是 log_{2}^{N}，
    查找的次数为 log_{2}^{N}，时间复杂度是 O(logN)
 */
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
