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
