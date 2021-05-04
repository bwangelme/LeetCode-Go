package l7

// 时间复杂度 log_10^(x)+1
// 空间复杂度 变量没有随着循环增加，可以认为是常数级, O(1)
func reverse(x int) int {
	var res int32
	for x != 0 {
		var positiveMax int32 = 214748364
		var negativeMin int32 = -214748364
		digit := x % 10
		// Int32Min -2147483647, Int32Max 2147483646
		// negativeMin = Int32Min - digit / 10
		// positiveMax = Int32Max - digit / 10
		// 根据 digit 的不同，Max 和 Min 会有 +-1 的浮动
		if digit >= 3 {
			negativeMin = -214748365
		}
		if digit <= -4 {
			positiveMax = 214748365
		}
		if res > positiveMax || res < negativeMin {
			return 0
		}
		res = res*10 + int32(digit)
		x = x / 10
	}

	return int(res)
}
