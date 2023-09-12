package sort

import (
	"fmt"
	"unicode"
)

//LetterGroup
/**
对字母进行分组

例如存在 "D a F B c A z" 字符串，需要将小写字母都放在大写字母前面，小写字母/大写字母内部的顺序无所谓
*/
func LetterGroup(chars []byte, n int) {
	charGroup(chars, n, func(b byte) bool {
		return unicode.IsLower(rune(b))
	})
}

//DigitLetterGroup
/**
对字母数字进行排序，小写字母放左边，数字放中间，大写字母放右边
*/
func DigitLetterGroup(chars []byte, n int) {
	a, _ := charGroup(chars, n, func(v byte) bool {
		return unicode.IsLower(rune(v)) || unicode.IsDigit(rune(v))
	})
	//fmt.Println("after lower || digit", outputChar(chars), outputChar(chars[0:a]), a)
	charGroup(chars[0:a], a, func(v byte) bool {
		return unicode.IsLower(rune(v))
	})
	//fmt.Println(outputChar(chars), chars[0:a]), a, n)
}

func charGroup(chars []byte, n int, leftFunc func(byte) bool) (int, int) {
	var (
		a = 0
		b = n - 1
	)

	for a < b {
		for leftFunc(chars[a]) && a < n {
			a++
		}
		for !leftFunc(chars[b]) && b >= 0 {
			b--
		}

		if a >= b {
			break
		}

		tmp := chars[a]
		chars[a] = chars[b]
		chars[b] = tmp
	}

	return a, b
}

func outputChar(arr []byte) (res string) {
	for i := 0; i < len(arr); i++ {
		res += fmt.Sprintf("%s ", string(arr[i]))
	}

	return res
}
