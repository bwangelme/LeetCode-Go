package l647

/*
## 题目
给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。

回文字符串 是正着读和倒过来读一样的字符串。

子字符串 是字符串中的由连续字符组成的一个序列。

具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。



示例 1：

输入：s = "abc"
输出：3
解释：三个回文子串: "a", "b", "c"
示例 2：

输入：s = "aaa"
输出：6
解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"


提示：

1 <= s.length <= 1000
s 由小写英文字母组成

## 解题思路

以 i 作为字符串中心，向左右两边延伸，统计回文字符串的个数

## 复杂度分析

时间复杂度: O(n^2)
空间复杂度: O(1)
*/

func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}

	var count = 0

	for i := 0; i < len(s); i++ {
		count += countPalindrome(s, i, i)
		count += countPalindrome(s, i, i+1)
	}

	return count
}

func countPalindrome(s string, start, end int) int {
	count := 0
	for start >= 0 && end < len(s) && s[start] == s[end] {
		count++
		start--
		end++
	}

	return count
}
