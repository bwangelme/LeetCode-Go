package l438

/*
## 题目

给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。



示例 1:

输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
 示例 2:

输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。


提示:

1 <= s.length, p.length <= 3 * 10^4
s 和 p 仅包含小写字母

## 解题思路

本题是 <567> 的变形题，因此解题思路和 567 相同。唯一区别是 567 中找到了子串之后直接返回 true, 这里是将子串起始索引存储起来，最后统一返回。

## 复杂度分析

和 <567> 相同

m == len(s)
n == len(p)
t == len(res)

时间复杂度: O(m+n)
空间复杂度: O(t)
t 表示 res 占用的空间，res 的大小无法根据 m, n 推导，因此用一个新的字母来表示

*/

func findAnagrams(s string, p string) []int {
	var res = make([]int, 0)

	if len(p) > len(s) {
		return res
	}

	var counts = make([]int, 26)
	for i, v := range p {
		counts[v-'a']++
		counts[s[i]-'a']--
	}

	if allZero(counts) {
		res = append(res, 0)
	}

	for i := len(p); i < len(s); i++ {
		counts[s[i]-'a']--
		counts[s[i-len(p)]-'a']++

		if allZero(counts) {
			res = append(res, i-len(p)+1)
		}
	}

	return res
}

func allZero(counts []int) bool {
	for _, c := range counts {
		if c != 0 {
			return false
		}
	}

	return true
}
