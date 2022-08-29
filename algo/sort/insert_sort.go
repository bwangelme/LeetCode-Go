package sort

import "github.com/bwangelme/LeetCode-Go/algo/ttypes"

/*
InsertSort 插入排序

## 算法思路

1. 从 1 开始遍历数组中每个元素
2. 从后往前遍历每个元素，找到第一个比当前元素小的元素，将其插入到该元素之后
	- 先挪位置，循环结束后再在 j 指示的索引中插入当前值

演示图: https://raw.githubusercontent.com/hustcc/JS-Sorting-Algorithm/master/res/insertionSort.gif
*/
func InsertSort[T ttypes.Number](arr []T) []T {
	res := make([]T, len(arr))
	copy(res, arr)

	for i := 1; i < len(res); i++ {
		j := i
		current := res[i]

		for j > 0 && res[j-1] > current {
			res[j] = res[j-1]
			j--
		}
		res[j] = current

	}

	return res
}

/*
InsertSortWithSentinel

使用了哨兵模式的插入排序

在数组头部插入一个空元素，然后将每次比较的元素放到 arr[0] 中。以此省略掉 for 循环中的一次比较
*/
func InsertSortWithSentinel[T ttypes.Number](arr []T) []T {
	res := make([]T, len(arr))
	copy(res, arr)

	res = append([]T{0}, res...)

	for i := 2; i < len(res); i++ {
		j := i
		current := res[i]

		res[0] = current
		// 和无哨兵相比，省略掉了 j > 0 这个比较条件
		for res[j-1] > current {
			res[j] = res[j-1]
			j--
		}

		res[j] = current

	}

	return res[1:]
}
