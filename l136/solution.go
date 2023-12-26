package l136

/*
## 分析

相同数字异或运算后是0, 因此将 nums 中所有数字异或运算后，得到的数字就是只出现一次的数字。
*/
func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res = res ^ num
	}

	return res
}
