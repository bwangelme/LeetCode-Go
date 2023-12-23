package l29

const MaxInt = 1<<31 - 1
const MinInt = -(1 << 31)

func IntAbs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}

func divideV2(dividend int, divisor int) int {
	if dividend == MinInt && divisor == -1 {
		return MaxInt
	}

	negative := 2
	if dividend > 0 {
		negative--
		dividend = -dividend
	}

	if divisor > 0 {
		negative--
		divisor = -divisor
	}

	res := divideCore(dividend, divisor)
	if negative == 1 {
		return -res
	} else {
		return res
	}
}

/*
## 解题思路

1. 先用 dividend 逐个减去 divisor (第一层 for 循环)
2. 除数以2的倍数不断增长，减少第一步减的次数 (第二层 for 循环)
3. value + value 不能溢出, 判断 value >= hafMinInt
*/
func divideCore(dividend int, divisor int) int {
	res := 0
	hafMinInt := -(1 << 30)

	for dividend <= divisor {
		value := divisor
		quotient := 1
		for value >= hafMinInt && dividend <= value+value {
			quotient += quotient
			value += value
		}

		res += quotient
		dividend -= value
	}

	return res
}

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		panic("divided by zero")
	}

	// 由于 dividend 和 divisor 都是 32 位有符号整数，所以溢出的情况只会是下面这一种情况。
	if dividend == MinInt && divisor == -1 {
		return MaxInt
	}

	var (
		dvd = IntAbs(dividend)
		dvs = IntAbs(divisor)
	)

	var sign = 1
	if (dividend > 0) != (divisor > 0) {
		sign = -1
	}

	var answer = 0
	for dvd >= dvs {
		var (
			tmp   = dvs
			count = 1
		)
		for tmp<<1 <= dvd {
			tmp <<= 1
			count <<= 1
		}
		dvd -= tmp
		answer += count
	}

	return answer * sign
}
