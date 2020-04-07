package l342

/*
## 题目描述

https://leetcode.com/problems/power-of-four/

__注意:__

1. 给定的整数参数是 `signed 32 bits`，即它有可能是负数
2. Follow up: Could you solve it without loops/recursion? 我们要给出一个不用循环或递归的解法

## 解题思路

这道题的 Related Topic 中给出的主题是 Bit Manipulation，我们要从4的二进制形式去考虑

4的N次方幂可以写成 4^N，根据指数运算的规则，可以进一步写成 2^{2N}。从这里我们就可以看出其二进制形式的规律
	1. 只有1位是1，其余都是0
	2. 为1的位其位置是个偶数
例如: 4 -> 1000, 0 -> 0001, 16 -> 0001 0000

找到了这个规律后，我们的解决方案就是逐个判断给定的数字是否满足上面的条件

0. num 必须是正数
1. num & (num-1) == 0，这个整数只有一位是1
2. num & 0x55555555 == 0，这个整数为1的位，它的位置索引是个偶数

## 复杂度分析

时间复杂度: O(1)
空间复杂度: O(1)
 */

func isPowerOfFour(num int) bool {
	if num < 0 {
		return false
	}

	// 判断 num 的二进制表示中是否只有1位是1
	if num & (num-1) != 0 {
		return false
	}

	// num 这一位1的位置是偶数 即，0,2,4,6,8,10
	if num & 0x55555555 != 0 {
		return true
	}

	return false
}
