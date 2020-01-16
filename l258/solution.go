package l258

// addDigits 计算 __数根__ 的程序
// https://www.bwangel.me/2019/04/09/leetcode-258/
func addDigits(num int) int {
	if num <= 0 {
		return 0
	} else {
		res := num % 9
		if res == 0 {
			return 9
		} else {
			return res
		}
	}
}
