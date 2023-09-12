package sort

//RadixSort
/**
基数排序

## 问题

对手机号进行排序

## 思路

从右边起，对每一位数字进行计数排序

## 分析

因为计数排序是稳定性排序，我们可以这样操作，如果计数排序是不稳定排序，就不能工作了。

## 复杂度分析

手机号一共有11位，每次计数排序的时间复杂度是 O(n), 得到整体时间复杂度是 O(11 * n)
**/
func RadixSort(phones []string, n int64) {
	if len(phones) == 0 {
		return
	}

	phoneLen := len(phones[0])
	for i := phoneLen - 1; i >= 0; i-- {
		PhoneCountingSort(phones, n, i)
	}
}

func PhoneCountingSort(arr []string, n int64, phoneIndex int) {
	var max int64 = 9 + 48
	var i int64

	var countingArr = make([]uint8, max+1)
	for i = 0; i < n; i++ {
		v := arr[i][phoneIndex]
		countingArr[v] += 1
	}

	for i = 1; i <= max; i++ {
		countingArr[i] = countingArr[i-1] + countingArr[i]
	}

	var resArr = make([]string, n)
	for i = n - 1; i >= 0; i-- {
		index := countingArr[arr[i][phoneIndex]] - 1
		resArr[index] = arr[i]
		countingArr[arr[i][phoneIndex]]--
	}

	for i := 0; i < len(resArr); i++ {
		arr[i] = resArr[i]
	}
}
