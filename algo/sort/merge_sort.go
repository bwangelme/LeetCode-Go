package sort

/*
## 分析

归并排序的思路

mergeSortC 对数组的子数组进行排序，当 len(子数组) == 1 的时候，这个数组就是有序的

merge 对有序的子数组进行合并, 就像交换需要创建 tmp 变量。merge 时 先创建 tmp 数组，将 left, right 两个子数组的有序结果放入 tmp 中，再将 tmp 更新到 arr 的对应位置

## 介绍

1. 无论数据源是什么，归并排序的时间复杂度固定是 O(nlogn)
2. 归并排序的空间复杂度是 O(1) ，因为每次只有一个函数会申请一个临时数组
3. 归并排序是稳定性排序，两个相同的元素，排序之后它们的顺序不会变

## 时间复杂度分析

T(1) = C 当数组长度为 1 时，时间复杂度是常数, 因为长度为1的数组本身是有序的，所以不需要排序
T(n) = 2 * T(n/2) + n

T(n) = 2*T(n/2) + n
     = 2*(2*T(n/4) + n/2) + n = 4*T(n/4) + 2*n
     = 4*(2*T(n/8) + n/4) + 2*n = 8*T(n/8) + 3*n
     = 8*(2*T(n/16) + n/8) + 3*n = 16*T(n/16) + 4*n
     ......
     = 2^k * T(n/2^k) + k * n
     ......
T(n) = 2^k * T(n/2^k) + k*n

当 T(n/2^k) == T(1) 时，得到 k = log_2{n}
得到

T(n) = n * T(1) + n * log_2{n}
T(n) = n * C + n * log_2{n}

用大O表示法，得到 归并排序的时间复杂度是 O(n*logn)

## 空间复杂度分析

merge 函数中每次都会申请临时数组 tmp, 因为每次 merge 函数执行完之后，tmp 数组就被释放掉了

所以归并排序申请的额外空间最大不会超过 N, 所以空间复杂度是 O(N)
*/

func MergeSort(arr []int64) []int64 {
	mergeSortC(arr, 0, len(arr)-1)

	return arr
}

func mergeSortC(arr []int64, start, end int) {
	if start >= end {
		return
	}

	q := (start + end) / 2
	mergeSortC(arr, start, q)
	mergeSortC(arr, q+1, end)
	merge(arr, start, q, q+1, end)
}

func merge(arr []int64, leftStart, leftEnd int, rightStart, rightEnd int) {
	var tmp = make([]int64, 0)
	var leftIdx, rightIdx = leftStart, rightStart

	for leftIdx <= leftEnd && rightIdx <= rightEnd {
		if arr[leftIdx] <= arr[rightIdx] {
			tmp = append(tmp, arr[leftIdx])
			leftIdx++
		} else {
			tmp = append(tmp, arr[rightIdx])
			rightIdx++
		}
	}

	var (
		remainStart = leftIdx
		remainEnd   = leftEnd
	)
	if leftIdx > leftEnd {
		remainStart = rightIdx
		remainEnd = rightEnd
	}

	for remainStart <= remainEnd {
		tmp = append(tmp, arr[remainStart])
		remainStart++
	}

	for idx := 0; idx <= (rightEnd - leftStart); idx++ {
		arr[leftStart+idx] = tmp[idx]
	}
}
