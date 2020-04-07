package l342

func isPowerOfFour(num int) bool {
	if num < 0 {
		return false
	}

	// 判断 num 的二进制表示中是否只有1位是1
	if num & (num -1) != 0 {
		return false
	}

	// num 这一位1的位置是偶数 即，0,2,4,6,8,10
	if num & 0x55555555 != 0 {
		return true
	}

	return false
}
