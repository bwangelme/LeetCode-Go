package l7

func reverse(x int) int {
	var res int32
	for x != 0 {
	    digit := x % 10
	    // 根据 digit 的不同，Max 和 Min 会有 +-1 的浮动
	    // 这里计算 positiveMax 和 negativeMin 还是利用 64 位正数来计算，只用32位整数计算的时候会溢出
		var positiveMax = int32((2147483646 - digit) / 10)
		var negativeMin = int32((-2147483647 - digit) / 10)
		if res > positiveMax || res < negativeMin {
			return 0
		}
		res = res*10 + int32(digit)
		x = x / 10
	}

	return int(res)
}
