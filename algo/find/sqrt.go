package find

import (
	"math"
)

//Sqrt
/**
使用二分查找实现开平方根，精确到小数点后六位

## 复杂度分析

时间复杂度是 O(logN)

它要求精确到小数点后6位，但整体趋势还是个对数形式的过程，所以时间复杂度是 O(logN)

没有申请新空间，空间复杂度是 O(1)
*/
func Sqrt(n float64) float64 {
	if n < 2 {
		return n
	}

	var (
		start float64 = 1
		end           = n / 2
		res   float64
		cnt   = 0
	)

	for start <= end {
		cnt++
		mid := start + (end-start)/2
		power := mid * mid
		if math.Abs(power-n) < 0.000001 {
			return mid
		}

		res = mid
		if power > n {
			end = mid
		} else {
			start = mid
		}
	}

	return res
}
