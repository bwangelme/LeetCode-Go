package l3_v2

/*
## 题目

给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。


提示：

0 <= s.length <= 5 * 10^4
s 由英文字母、数字、符号和空格组成

## 解题思路

counts 是保存字符出现次数的 map
j, i 是指向子串的左右指针

i 不断向右移动，并更新 counts，如果遇到重复字符 (counts[s[i]]==2)
则向右移动 j ，直到把重复字符删除掉。

每次移动到没有重复字符的时候，更新一下 maxLen(i-j)

## 复杂度分析

n 表示 len(s)
时间复杂度: O(n) (i 和 j 最多移动 n 次)
空间复杂度：O(1) (counts 是个固定长度的数组)
*/

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var counts = make([]int, 256)
	var (
		countDup = 0
		// 因为 j 是先++ 再索引，因此从 -1 开始
		j      = -1
		maxLen = 1
	)

	for i, v := range s {
		counts[v]++
		if counts[v] == 2 {
			countDup++
		}

		for countDup > 0 {
			j++
			counts[s[j]]--
			if counts[s[j]] == 1 {
				countDup--
			}
		}

		maxLen = max(maxLen, i-j)
	}

	return maxLen
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
