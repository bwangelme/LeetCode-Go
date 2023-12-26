package l137

/*
## 解题思路

输入 2 2 3 2, 它们的二进制表示是

00 00 00 10
00 00 00 10
00 00 00 11
00 00 00 10

将所有数字的二进制表示求和，得到

00 00 00 41

三个相同的数，包含1的位，求和得到的值一定是3的倍数，不是1的倍数的位，就是只出现一次的数字的二进制位

将上述结果每一位对 3 取余，得到

00 00 00 11

即得到只出现一次的数字 3

## 复杂度分析

n 表示 nums 的长度

- 空间复杂度

bitSum 是固定大小的32位，因此空间复杂度是 O(1)

- 时间复杂度

O(N * 32) + O(32) = O(N)
*/
func singleNumber(nums []int) int {
	var bitSum = make([]uint8, 32)
	for _, num := range nums {
		for i := 0; i < 32; i++ {
			bitSum[i] += uint8(num >> (31 - i) & 1)
		}
	}

	var res = int32(0)
	for i := 0; i < 32; i++ {
		res = res<<1 + int32(bitSum[i]%3)
	}

	return int(res)
}
