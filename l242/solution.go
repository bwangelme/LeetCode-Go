package l242

/*
## 题目

给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false

提示:

1 <= s.length, t.length <= 5 * 10^4
s 和 t 仅包含小写字母


进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

## 复杂度分析

N == utf8_len(s)
M == s 中不同字符的种类个数
时间复杂度: O(N), for 循环了两次
空间复杂度: O(M)
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var counts = make(map[rune]int)
	for _, ch := range s {
		counts[ch]++
	}

	for _, ch := range t {
		v, ok := counts[ch]
		if !ok {
			return false
		}
		if v <= 0 {
			return false
		}
		counts[ch]--
	}

	return true
}
