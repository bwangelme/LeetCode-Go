package l820

/*
## 题目描述
单词数组 words 的 有效编码 由任意助记字符串 s 和下标数组 indices 组成，且满足：

words.length == indices.length
助记字符串 s 以 '#' 字符结尾
对于每个下标 indices[i] ，s 的一个从 indices[i] 开始、到下一个 '#' 字符结束（但不包括 '#'）的 子字符串 恰好与 words[i] 相等
给你一个单词数组 words ，返回成功对 words 进行编码的最小助记字符串 s 的长度 。



示例 1：

输入：words = ["time", "me", "bell"]
输出：10
解释：一组有效编码为 s = "time#bell#" 和 indices = [0, 2, 5] 。
words[0] = "time" ，s 开始于 indices[0] = 0 到下一个 '#' 结束的子字符串，如加粗部分所示 "time#bell#"
words[1] = "me" ，s 开始于 indices[1] = 2 到下一个 '#' 结束的子字符串，如加粗部分所示 "time#bell#"
words[2] = "bell" ，s 开始于 indices[2] = 5 到下一个 '#' 结束的子字符串，如加粗部分所示 "time#bell#"
示例 2：

输入：words = ["t"]
输出：2
解释：一组有效编码为 s = "t#" 和 indices = [0] 。


提示：

1 <= words.length <= 2000
1 <= words[i].length <= 7
words[i] 仅由小写字母组成

## 解题思路

1. 将单词反转后，合并后缀相同的单词就变成了合并前缀相同的单词，就可以使用前缀树来解
2. 将所有单词构造成一颗前缀树，统计从根节点到每个叶子节点的长度和，就是拼接的字符串的长度(因此每个单词后面有个 #，因此统计时需要从 root=1 开始计数)

## 复杂度分析

N = 单词的个数
M = 每个单词的平均长度

### 空间复杂度

空间复杂度为构造的前缀树的复杂度，最坏情况下为 O(N * M)

### 时间复杂度

1. 构造前缀树的时间复杂度 O(N * M),针对每个单词都运行一遍构造逻辑
2. dfs的时间复杂度，将每个叶子节点都遍历了一遍，因此最坏时间复杂度是 O(N * M)
*/

type TrieNode struct {
	Children [26]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children: [26]*TrieNode{},
	}
}

func minimumLengthEncoding(words []string) int {
	root := buildTrie(words)
	var total = 0
	// 拼好的字符串中，每个单词后面都有一个#
	// 正确的单词长度是 len(word)+1, 正好等于从 root 开始统计到叶子节点,因此 length 从 1 开始
	dfs(root, 1, &total)
	return total
}

func buildTrie(words []string) *TrieNode {
	root := NewTrieNode()
	for _, word := range words {
		node := root
		chars := []byte(word)

		for i := len(chars) - 1; i >= 0; i-- {
			idx := chars[i] - 'a'
			if node.Children[idx] == nil {
				node.Children[idx] = NewTrieNode()
			}
			node = node.Children[idx]
		}
	}

	return root
}

func dfs(root *TrieNode, length int, total *int) {
	var isLeaf = true
	for _, child := range root.Children {
		if child != nil {
			isLeaf = false
			dfs(child, length+1, total)
		}
	}

	if isLeaf {
		*total = *total + length
	}
}
