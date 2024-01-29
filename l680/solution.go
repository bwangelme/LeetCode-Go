package l680

/*
## 题目

给你一个字符串 s，最多 可以从中删除一个字符。

请你判断 s 是否能成为回文字符串：如果能，返回 true ；否则，返回 false 。



示例 1：

输入：s = "aba"
输出：true
示例 2：

输入：s = "abca"
输出：true
解释：你可以删除字符 'c' 。
示例 3：

输入：s = "abc"
输出：false


提示：

1 <= s.length <= 10^5
s 由小写英文字母组成

## 复杂度分析

n == len(s)
时间复杂度: O(n), 因为整个过程遍历了一遍 s, 和部分的 s, 总遍历次数小于 2n, 因此时间复杂度是 O(n)
空间复杂度: O(n), 申请了 chars 这个额外的数组
*/

func validPalindrome(s string) bool {
	var (
		start = 0
		end   = len(s) - 1
		chars = []rune(s)
	)

	for start < len(s)/2 {
		if chars[start] != chars[end] {
			break
		}

		start++
		end--
	}

	return start >= len(s)/2 || isPalindrome(chars, start+1, end) || isPalindrome(chars, start, end-1)

}

func isPalindrome(chars []rune, start, end int) bool {
	for start < end {
		if chars[start] != chars[end] {
			break
		}
		start++
		end--
	}

	return start >= end
}
