package l208

/*
## 题目描述
Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补全和拼写检查。

请你实现 Trie 类：

Trie() 初始化前缀树对象。
void insert(String word) 向前缀树中插入字符串 word 。
boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。

## 结题思路

前缀树是一个 N 叉树，每个节点是单词的一个字符
*/

type TrieNode struct {
	Children [26]*TrieNode
	IsWord   bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		IsWord:   false,
		Children: [26]*TrieNode{},
	}
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{},
	}
}

func (this *Trie) Insert(word string) {
	runes := []rune(word)
	node := this.root

	for _, char := range runes {
		idx := char - 'a'
		v := node.Children[idx]
		if v == nil {
			node.Children[idx] = NewTrieNode()
		}

		node = node.Children[idx]
	}
	node.IsWord = true
}

func (this *Trie) Search(word string) bool {
	runes := []rune(word)

	node := this.root
	for _, char := range runes {
		idx := char - 'a'
		v := node.Children[idx]
		if v == nil {
			return false
		}
		node = v
	}

	return node.IsWord
}

func (this *Trie) StartsWith(prefix string) bool {
	runes := []rune(prefix)

	node := this.root
	for _, char := range runes {
		idx := char - 'a'
		v := node.Children[idx]
		if v == nil {
			return false
		}
		node = v
	}

	return true
}
