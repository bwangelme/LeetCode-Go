package l258

//addDigits 计算 __数根__ 的程序
// 维基百科链接： https://www.wikiwand.com/en/Digital_root#/Abstract_multiplication_of_digital_roots
func addDigits(num int) int {
	if num == 0 {
		return 0
	} else {
		return num - 9 * ((num -1) / 9)
	}
}
