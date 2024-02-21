package l14

/*
## 题目描述

编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。



示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。


提示：

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] 仅由小写英文字母组成

## 解题思路

挨个遍历每个单词的同一个字母，看字母是否相同

## 复杂度分析

n = len(strs)
m = max(len(str) for str in strs)

时间复杂度， O(n * 匹配的长度), 最坏为 O(n * m)
空间复杂度: O(1)
*/

func longestCommonPrefix(strs []string) string {
	var n = len(strs)
	if n == 0 {
		return ""
	}

	var i = 0

Outer:
	for {
		if i >= len(strs[0]) {
			break Outer
		}
		var char = strs[0][i]

		for j := 1; j < n; j++ {
			if i >= len(strs[j]) {
				break Outer
			}

			if strs[j][i] != char {
				break Outer
			}
		}
		i++
	}

	return strs[0][:i]
}
