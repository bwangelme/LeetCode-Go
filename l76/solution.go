package l76

import (
	"math"
)

/*
## 题目

给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：

对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。

示例 1：

输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
示例 2：

输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 是最小覆盖子串。
示例 3:

输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。

提示：

m == s.length
n == t.length
1 <= m, n <= 10^5
s 和 t 由英文字母组成

## 注意

i, v := range t

v 由于是按照 unicode 码点遍历的，v 的类型是 int32

v := t[3]

是按照索引找到的字符，v 的类型是 uint8

## 解题思路

1. 用 countMap 存储在 t 中出现的字符的个数
2. count 表示在 t 中出现，还未在子串 s[start:end] 中出现的字符的个数
3. 增加 end ，扩大 s[start:end], 让 count == 0, 找到包含 t 中所有字符的子串
4. 增加 start, 减少 s[start:end]，找最短子串, 如果有在 t 中存在的字符++导致 countMap[char] == 1, 则 count++，如果 start 增加导致 count > 0, 那么重新增加 end, 让 count--
5. end 遍历完整个 s, 在这个过程中找到了所有符合要求的子串，minLen,minStart,minEnd 记录了最短子串

## 复杂度分析

m == len(t)
n == len(s)

时间复杂度

开始遍历 t 的复杂度是 O(m)
start, end 一起遍历了 s, 时间复杂度是 O(n)

因为 m 一定小于 n， 因此 O(m+n) 可以写成 O(n)

空间复杂度
---

只申请了额外空间 countMap, 因为 s, t 中都是英文字母, countMap 的长度不超过 256，因此空间复杂度是 O(1)
*/
func minWindow(s string, t string) string {
	var (
		// 注意: len(t) 最大是 10^5, 应该考虑 countMap 的 value 最大能够放下 10^5
		countMap = make(map[uint8]int)

		start = 0
		end   = 0

		minLen   = math.MaxInt
		minStart = 0
		minEnd   = 0
	)

	// 将 t 中出现的所有字符在 countMap 中保存出现次数
	for i, _ := range t {
		v := t[i]
		_, ok := countMap[v]
		if !ok {
			countMap[v] = 1
		} else {
			countMap[v] += 1
		}
	}

	// count 表示在 t 中出现，还没在子串 s[start:end] 中出现的字符的个数
	// count == 0 时表示我们已经找到了包含 t 中所有字符的子串，此时需要调整 start, 减少子串长度来找最短子串
	var count = len(countMap)

	for end < len(s) || count == 0 && end == len(s) {
		if count > 0 {
			_, ok := countMap[s[end]]
			if ok {
				countMap[s[end]]--
				if countMap[s[end]] == 0 {
					count--
				}
			}
			end++
		} else {
			if (end - start) < minLen {
				minLen = end - start
				minStart = start
				minEnd = end
			}
			_, ok := countMap[s[start]]
			if ok {
				countMap[s[start]]++
				if countMap[s[start]] == 1 {
					count++
				}
			}
			start++
		}
	}

	if minLen < math.MaxInt {
		return s[minStart:minEnd]
	} else {
		return ""
	}
}
