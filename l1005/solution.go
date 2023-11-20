package l1005

/**
## 题目

给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组：

    选择某个下标 i 并将 nums[i] 替换为 -nums[i] 。

重复这个过程恰好 k 次。可以多次选择同一个下标 i 。

以这种方式修改数组后，返回数组 可能的最大和 。

提示：

    1 <= nums.length <= 10^4
    -100 <= nums[i] <= 100
    1 <= k <= 10^4

## 解题思路

将数组的值和次数，映射到一个数组中，这样寻找最小的数就不用排序了，直接遍历数组即可

每次取相反数都是取的最小的数

## 复杂度分析

N 表示 nums 数组的长度
K 表示反转的次数

### 时间复杂度

O(N) + O(200) + O(K * C) = O(N) + O(K)

### 空间复杂度

申请了额外的数组，空间复杂度是 O(200) = O(C)
*/

func largestSumAfterKNegations(nums []int, k int) int {
	var countMap = make([]int, 201)
	for i := 0; i < len(nums); i++ {
		countMap[nums[i]+100] += 1
	}

	var minNum = 0
	for k > 0 {
		// 最坏情况下，循环次数是 200，只有第一次最坏，后续循环次数就会小很多
		for countMap[minNum] == 0 {
			minNum += 1
			continue
		}

		countMap[minNum] -= 1
		countMap[200-minNum] += 1

		if minNum > 100 {
			minNum = 200 - minNum
		}
		k--
	}

	var sum = 0
	for j := 0; j < len(countMap); j++ {
		sum += countMap[j] * (j - 100)
	}

	return sum
}
