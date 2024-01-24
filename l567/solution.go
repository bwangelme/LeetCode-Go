package l567

/*
## 题目

给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。

换句话说，s1 的排列之一是 s2 的 子串 。

示例 1：

输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").
示例 2：

输入：s1= "ab" s2 = "eidboaoo"
输出：false


提示：

1 <= s1.length, s2.length <= 10^4
s1 和 s2 仅包含小写字母

## 解题思路

s1 的排列之一是 s2 的子串，我们只要统计 s1 的长度以及各个字母出现的频率，s2 中相同长度的子传有同频率的字母，那么 s2 中就有是 s1 排列之一的子串

counts 是一个 map, 用于统计每个字母出现的顺序。

先遍历 s1, 将 counts 中对应的字母 +1
再以固定长度遍历 S2, 将 counts 中对应的字母 -1, 如果遍历完之后，counts 中全部都是 0，说明遍历的这部分固定长度的子串是 S1 的排列之一。

## 复杂度分析

m 表示 s1 长度
n 表示 s2 长度

由于遍历了 S1 和 S2, 因此时间复杂度是 O(m+n)

注意，尽管 allZero 是个 for 循环，但由于 counts 是个常数，因此 allZero 的复杂度是 O(1)

由于新申请的数组 counts 是个常数，因此空间复杂度是 O(1)


*/

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var counts = make([]int, 26)

	// 遍历 s2[0:len(s1)] 的部分
	for i, v := range s1 {
		counts[v-'a']++
		counts[s2[i]-'a']--
	}

	if allZero(counts) {
		return true
	}

	// 遍历 s2[len(s1):len(s2)] 的部分
	// 向前移动的过程中，字串相当于是删除了 s2[i-len(s1)], 加上了 s2[i]
	// 因此 counts 也可以做同样处理，删除掉的字母 +1, 新增的字母 -1
	for i := len(s1); i < len(s2); i++ {
		counts[s2[i]-'a']--
		counts[s2[i-len(s1)]-'a']++
		if allZero(counts) {
			return true
		}
	}

	return false
}

func allZero(counts []int) bool {
	for _, c := range counts {
		if c != 0 {
			return false
		}
	}

	return true
}
