package l49

/*
## 题目

给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:

输入: strs = [""]
输出: [[""]]
示例 3:

输入: strs = ["a"]
输出: [["a"]]


提示：

1 <= strs.length <= 10^4
0 <= strs[i].length <= 100
strs[i] 仅包含小写字母

*/

/*
## 解题思路

key := [26]int{}
key 当作 map 的 key, key 实现了 == 和 != 方法，而且变位词的 key 相等

遍历所有单词，就可以将单词按照 key 分组了

## 复杂度分析

m == 每个单词的平均长度
n == 单词个数
k == 结果的分组个数

时间复杂度

O(m * n),

每个遍历每个单词计算 key 是 O(m), 一共有 n 个单词

空间复杂度

O(m * n + k), 主要分析 groups map 占用的空间

k 表示 map 的 key 占用的空间
m * n 表示 map 的 value 占用的空间
*/
func groupAnagrams(strs []string) [][]string {
	var (
		groups = make(map[[26]int][]string)
	)
	for _, str := range strs {
		key := [26]int{}
		for _, ch := range str {
			idx := ch - 'a'
			key[idx]++
		}
		groups[key] = append(groups[key], str)
	}
	res := make([][]string, 0)
	for _, v := range groups {
		res = append(res, v)
	}
	return res
}

/*
## 解题思路

每个字母映射到一个质数上，如果两个单词的字母乘积相同，那么这两个单词是变位词

## 复杂度分析

m == 每个单词的平均长度
n == 单词个数
k == 结果的分组个数

时间复杂度

O(m * n), 每个单词执行了 m 次相乘，共有 n 个单词

空间复杂度

O(m * n + k), 主要分析 groups map 占用的空间

k 表示 map 的 key 占用的空间
m * n 表示 map 的 value 占用的空间
*/
func groupAnagramsV1(strs []string) [][]string {
	var (
		letterWeight = []int{
			2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101,
		}
		groups = make(map[uint64][]string)
	)

	for _, str := range strs {
		var wordWeight uint64 = 1
		for i := 0; i < len(str); i++ {
			lWeight := uint64(letterWeight[str[i]-'a'])
			wordWeight *= lWeight
		}
		v, ok := groups[wordWeight]
		if !ok {
			v = make([]string, 0)
		}
		v = append(v, str)
		groups[wordWeight] = v
	}

	res := make([][]string, 0)
	for _, v := range groups {
		res = append(res, v)
	}
	return res
}
