package sort

// QuickSort 快速排序
/**
## 思路

它的思路是，给定一个数组 arr

- partition 函数找到支撑点 pivot 的索引 p，保证 p 左边的数字都 < pivot，右边的数字都 >= pivot
	- 在  partition 函数中巧妙地实现了原地排序，没有申请新的数组空间
- 然后再用分治的思想，分别将 p 左右两边的数组再进行排序

快排和归并排序的过程正好是相反的，

- 归并排序是自底向上，先将长度为1的数组进行排序，再逐级往上
- 快排是自顶向下，先将数组分成两份，保证 all of left < pivot < all of right, 再对左右两个数组进行排序

## 时间复杂度

大部分情况下，时间复杂度都是 O(n*logn), 最坏情况下时间复杂度是 O(n^2)

### 最坏情况

例如 1, 3, 5, 6, 8 这个有序数组，每次分区只得到一个数组, 左边是 n-1, 右边是 0

T(1) = C
T(n) = T(n-1) + n
     = T(n-2) + n - 1 + n
     ....
     = T(n-n-1) + 2 + 3 + 4 ... + n

T(n) = C + 2 + 3 +4 + .. n = C + (1+n)*n/2 - 1

用大 O 表示法就是 O(n^2)

### 最好情况

每次分区，得到两个大小相同的数组，那么它的复杂度就和归并排序完全相同，就是 O(n*log_2{n})

## 空间复杂度

快速排序没有申请新数组空间，它的空间复杂度是 O(1)

## 介绍

1. 快速排序不是稳定排序
给定 6, 8, 7, 6, 3, 5, 9, 4 这个数组，在第一次调用 parition 之后，6 和 6 的相对位置已经发生了变化
**/
func QuickSort(arr []int64) []int64 {
	n := len(arr) - 1
	quickSortC(arr, 0, n)

	return arr
}

func quickSortC(arr []int64, p, r int) {
	if p >= r {
		return
	}

	q := partition(arr, p, r)
	quickSortC(arr, p, q-1)
	quickSortC(arr, q+1, r)

}

func partition(arr []int64, p, r int) int {
	pivot := arr[r]
	i := p
	for j := p; j <= r-1; j++ {
		if arr[j] < pivot {
			if i != j {
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}

			i++
		}
	}

	tmp := arr[i]
	arr[i] = arr[r]
	arr[r] = tmp

	return i
}
