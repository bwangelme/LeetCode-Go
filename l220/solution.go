package l220

func getBucketId(num int, bucketSize int) int {
	if num >= 0 {
		return num / bucketSize
	} else {
		return (num+1)/bucketSize - 1
	}
}

/*
## 题目描述

给你一个整数数组 nums 和两个整数 indexDiff 和 valueDiff 。

找出满足下述条件的下标对 (i, j)：

i != j,
abs(i - j) <= indexDiff
abs(nums[i] - nums[j]) <= valueDiff
如果存在，返回 true ；否则，返回 false 。

## 解题思路

我们将输入的数字放到若干个桶中， getBucketsId 函数

[-valueDiff-1, -1] 为 -1 号桶
[0, valueDiff] 为 0 号桶
[valueDiff+1, 2*valueDiff+1] 为 1 号桶

以此类推，每个桶的大小是 valueDiff+1

for 循环遍历数组的每个数 num，下标为 i containsNearbyAlmostDuplicate 函数

	如果 num 所在的桶中已经有数字了，这两个数字满足差的绝对值 <= valueDiff，那么 return true
	如果 num 所在桶的前一个桶有数字 numP，满足 num - numP <= valueDiff，那么 return true
	如果 num 所在桶的后一个桶有数字 numN，满足 numN - num <= valueDiff，那么 return true

	将 num 放入它的桶中
	判断如果 i >= indexDiff
		此时所有桶中的已经存在不满足 indexDiff 的数字了，将 nums[i-indexDiff] 从它的桶中移出去

## 复杂度分析

# N 表示 nums 数组的长度

空间复杂度，桶的个数最多为 indexDiff 个，因此空间复杂度是 O(indexDiff)
时间复杂度，遍历每个数字中，从桶中取，放数字的时间复杂度都是 O(1)，因此总的时间复杂度是 O(N)
*/
func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	var buckets = make(map[int]int, 0)
	for i, num := range nums {
		// 在循环中的数字一定满足 /i-j/ <= indexDiff
		// 因为循环结尾会从桶中删除多余的数字
		bid := getBucketId(num, valueDiff+1)

		// 同一个桶里有两个数字，它们
		_, ok := buckets[bid]
		if ok {
			return true
		}

		v, ok := buckets[bid-1]
		// 前一个桶中存在和 num 的差的绝对值 <= valueDiff 的数字
		if ok && v+valueDiff >= num {
			return true
		}

		// 后一个桶中存在和 num 的差的绝对值 <= valueDiff 的数字
		v, ok = buckets[bid+1]
		if ok && v-valueDiff <= num {
			return true
		}

		// 如果一个桶中存在两个数字，前面的判断条件已经让它们返回 true 了
		// 因此一个桶只要存一个数字就好
		buckets[bid] = num
		if i >= indexDiff {
			delete(buckets, getBucketId(nums[i-indexDiff], valueDiff+1))
		}
	}

	return false
}
