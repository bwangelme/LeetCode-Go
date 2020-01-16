package l258

// addDigits 计算 __数根__ 的程序
// https://www.bwangel.me/2019/04/09/leetcode-258/
func addDigits(num int) int {
	if num <= 0 {
		// 非自然数及0的情况
		return 0
	} else {
		res := num % 9
		if res == 0 {
			// 模9等于0时，其数根为9
			return 9
		} else {
			// 自然数的数根为其模9的余数
			return res
		}
	}
}
