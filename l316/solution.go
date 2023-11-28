package l316

/*
## 题目

给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。

提示：

    1 <= s.length <= 104
    s 由小写英文字母组成


## 解题思路

1. 用 remainCounter 保存每个字母出现的次数
2. 用 res 保存结果

遍历字符串，如果当前字符 c 没在 res 中出现，尝试将 c 放入 res

将 c 往 res 里面放的时候，检查 res 中的所有元素，将比 c 大，而且后续还会出现(remainCounter)的字符丢弃

## 复杂度分析

n 表示字符串的长度

### 空间复杂度

o(n * 3)

### 时间复杂度

o(n) + o(n^2)

每次判断 res 中的字符可以弹出的时候，最坏情况下都会弹出，因此时间复杂度是 O(n^2)

*/

func removeDuplicateLetters(s string) string {
	var (
		res           []rune
		resExist      = make(map[rune]struct{})
		remainCounter = make(map[rune]int)
	)

	for _, c := range s {
		remainCounter[c] += 1
	}

	for _, c := range s {
		_, exist := resExist[c]
		if !exist {
			last := len(res) - 1
			for last >= 0 && c < res[last] && remainCounter[res[last]] > 0 {
				delete(resExist, res[last])
				res = res[0:last]
				last = len(res) - 1
			}
			res = append(res, c)
			resExist[c] = struct{}{}
		}
		remainCounter[c] -= 1
	}

	return string(res)
}
