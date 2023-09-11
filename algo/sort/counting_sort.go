package sort

//CountingSort
/**

## 适用场景

1. 对非负整数进行计数排序, 或者可以将其他数字转换成非负证书
2. 数组中的值在一个小范围内波动，例如年龄(0-200)，考试分数(0-750)

## 思路

1. 创建 countingArr, countingArr 中的数字表示 <= index 的计数，
	例如 countingArr[5] = 8, 表示整个数组中 <= 5 的数字有8个

2. 根据 countingArr 排序
	2.1 逆序遍历数组, for v in reverse(arr), 创建新数组 newArr 用于放排序结果
	2.2 countingArr[v] == t ， 表示小于等于 v 的数字有 t 个，那么 v 应该放在数组的第 t 位， 即 newArr[t-1] = i
	2.3 v 放入到新数组后，对应的计数值也 -1, countingArr[v]--
	2.4 遍历完原始数组 arr 后，所有的 arr 都按合适顺序放入到了 newArr 中

## 复杂度分析

程序都是以 N 为单位进行 for 循环，时间复杂度一眼就看出是 O(N)

空间复杂度也是 O(N)
*/
func CountingSort(arr []int64, n int64) []int64 {
	var max int64 = -1
	var i int64

	// 1. 查找 max, 创建数组 countingArr
	for i = 0; i < n; i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}

	var countingArr = make([]int64, max+1)
	// 2. 统计每个数字的个数, 并将统计数字累加
	for i = 0; i < n; i++ {
		countingArr[arr[i]] += 1
	}

	for i = 1; i <= max; i++ {
		countingArr[i] = countingArr[i-1] + countingArr[i]
	}

	// 3. 创建 resArr ，用于存储结果
	var resArr = make([]int64, n)
	// 4. 根据 countingArr, 将结果排序
	for i = n - 1; i >= 0; i-- {
		index := countingArr[arr[i]] - 1
		resArr[index] = arr[i]
		countingArr[arr[i]]--
	}

	return resArr
}
