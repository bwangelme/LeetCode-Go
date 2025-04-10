package l676

/*
## 题目描述

设计一个使用单词列表进行初始化的数据结构，单词列表中的单词 互不相同 。 如果给出一个单词，请判定能否只将这个单词中一个字母换成另一个字母，使得所形成的新单词存在于你构建的字典中。

实现 MagicDictionary 类：

MagicDictionary() 初始化对象
void buildDict(String[] dictionary) 使用字符串数组 dictionary 设定该数据结构，dictionary 中的字符串互不相同
bool search(String searchWord) 给定一个字符串 searchWord ，判定能否只将字符串中 一个 字母换成另一个字母，使得所形成的新字符串能够与字典中的任一字符串匹配。如果可以，返回 true ；否则，返回 false 。

## 复杂度分析

时间复杂度

buildDict
N == len(dictionary) 单词的个数
假设 dictionary 中单词的平均长度是 M
buildDict 的平均时间复杂度是 O(M * N)

Search

L == len(searchWord)
最坏时间复杂度是 O(L)，最多会查找 L 次

最坏空间复杂度, O(M * N)，输入的N个单词，彼此没有复用，因此最坏空间复杂度是 O(M * N)，每个字母都申请了 TrieNode
*/

import (
	"github.com/bwangelme/LeetCode-Go/lt"
)

type MagicDictionary struct {
	trie *lt.Trie
}

func Constructor() MagicDictionary {
	return MagicDictionary{
		trie: lt.NewTrie(),
	}
}

func (m *MagicDictionary) BuildDict(dictionary []string) {
	for _, d := range dictionary {
		m.trie.Insert(d)
	}
}

func (m *MagicDictionary) Search(searchWord string) bool {
	runes := []rune(searchWord)
	return m.dfs(m.trie.Root, runes, 0, 0)
}

func (m *MagicDictionary) dfs(root *lt.TrieNode, runes []rune, i int, edit int) bool {
	if root == nil {
		return false
	}
	if root.IsWord && i == len(runes) && edit == 1 {
		return true
	}

	if i < len(runes) && edit <= 1 {
		found := false
		for j := 0; j < 26 && !found; j++ {
			searchIdx := int(runes[i] - 'a')
			next := edit
			if j != searchIdx {
				next = edit + 1
			}
			found = m.dfs(root.Children[j], runes, i+1, next)
		}
		return found
	}
	return false
}
