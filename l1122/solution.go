package l1122

/*
## 题目描述

给你两个数组，arr1 和 arr2，arr2 中的元素各不相同，arr2 中的每个元素都出现在 arr1 中。

对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾。



示例 1：

输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
输出：[2,2,2,1,4,3,3,9,6,7,19]
示例  2:

输入：arr1 = [28,6,22,8,44,17], arr2 = [22,28,8,6]
输出：[22,28,8,6,17,44]


提示：

1 <= arr1.length, arr2.length <= 1000
0 <= arr1[i], arr2[i] <= 1000
arr2 中的元素 arr2[i]  各不相同
arr2 中的每个元素 arr2[i] 都出现在 arr1 中

## 解题思路

本题使用计数排序的解法

1. 将 arr1 中所有出现的数字的次数放到 counts 中
2. 顺序遍历 arr2, 如果 num 在 counts 中出现过，按出现数字添加到 arr1 中
3. 遍历 counts, 将未出现的数字按次数添加到 arr1 中

## 复杂度分析

N = len(arr1)
M = len(arr2)

空间复杂度，counts 数组的大小是 arr 中数字的范围，根据题目可知是 O(1001)

时间复杂度

遍历了三个数组，arr1, arr2 和 counts, 时间复杂度是 O(N) + O(M) + O(1001) = O(M+N)


*/

func relativeSortArray(arr1 []int, arr2 []int) []int {
	var counts = [1001]int{}
	for _, num := range arr1 {
		counts[num]++
	}

	var i = 0
	for _, num := range arr2 {
		v := counts[num]
		for v > 0 {
			arr1[i] = num
			i++
			v--
		}
		counts[num] = v
	}

	for num, count := range counts {
		for count > 0 {
			arr1[i] = num
			i++
			count--
		}
	}

	return arr1
}
