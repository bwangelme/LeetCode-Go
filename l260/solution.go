package l260

/*
## 题目描述

给你一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。

你必须设计并实现线性时间复杂度的算法且仅使用常量额外空间来解决此问题。

提示

- 2 <= nums.length <= 3 * 10^4
- -2^31 <= nums[i] <= 2^31 - 1
- 除两个只出现一次的整数外，nums 中的其他数字都出现两次

## 解题思路

1. 将所有数字异或，得到了两个不同数字的异或结果
2. 从左往右找异或结果中不同的第一位的索引 index
3. 遍历所有结果，将 index == 0 的所有数字异或放到 res[0]，将 index == 1 的所有数字异或放到 res[1]

因为其他数字都相同，所以 res[0] 中一定只有一个数字, 其他的数字因为相同被异或成0了，res[1] 也是同样的结果

## 复杂度分析

- N 表示 nums 的长度

- 时间复杂度

O(2N) == O(N)

- 空间复杂度: O(1)

因为没有申请任何与 N 有关的数据结构
*/
func singleNumber(nums []int) []int {
	// 得到只出现两次的数字的异或结果
	var res = 0
	for _, num := range nums {
		res = res ^ num
	}

	// 从右向左查找异或结果中第一个不是0的位
	var index = 0
	for ; index < 32; index++ {
		if (res>>index)&1 == 1 {
			break
		}
	}

	var result = make([]int, 2)
	for _, num := range nums {
		if (num>>index)&1 == 1 {
			result[0] ^= num
		} else {
			result[1] ^= num
		}
	}

	return result
}
