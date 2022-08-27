package l541

import (
	"strings"
)

//reverseStr
// 题目解析
// 通俗一点说，每隔k个反转k个，末尾不够k个时全部反转；
func reverseStr(s string, k int) string {
	var res strings.Builder

	for i := 0; i*k < len(s); i++ {
		start := i * k
		end := (i + 1) * k

		if end > len(s) {
			end = len(s)
		}

		var sub = s[start:end]
		if i%2 == 0 {
			sub = reverse(sub)
		}

		res.WriteString(sub)
	}
	return res.String()
}

func reverse(s string) string {
	// sArr 是新建的数组 和s 没有关系
	var sArr = []rune(s)

	for i, j := 0, len(sArr)-1; i < j; i, j = i+1, j-1 {
		var tmp = sArr[i]
		sArr[i] = sArr[j]
		sArr[j] = tmp
	}

	return string(sArr)
}
