package l648

import "strings"

/*
## 题目描述

在英语中，我们有一个叫做 词根(root) 的概念，可以词根 后面 添加其他一些词组成另一个较长的单词——我们称这个词为 衍生词 (derivative)。例如，词根 help，跟随着 继承词 "ful"，可以形成新的单词 "helpful"。

现在，给定一个由许多 词根 组成的词典 dictionary 和一个用空格分隔单词形成的句子 sentence。你需要将句子中的所有 衍生词 用 词根 替换掉。如果 衍生词 有许多可以形成它的 词根，则用 最短 的 词根 替换它。

你需要输出替换之后的句子。

## 结题思路

1. 使用所有的词根字典构造一颗前缀树
2. 将句子中的每个单词输入，调用 FindPrefix

FindPrefix 的逻辑是
查看输入的单词，它的前缀是否匹配前缀树中的一个单词，如果匹配，返回匹配的前缀，如果不匹配，返回空
匹配到前缀树中的单词，立刻返回，这样查到的前缀就是最短的

## 复杂度分析

len(directory) = M
directory 中单词的平均长度 = K
len(split(sentence, " ")) = N

## 时间复杂度

构造前缀树，时间复杂度 = O(M * K)
针对 sentence 中的每个单词，执行 FindPrefix = O(N * K)

总的时间复杂度 = O((M+N) * K)

## 空间复杂度

本程序的额外申请的空间就是前缀树

最坏情况下，directory 中的单词没有复用的前缀，空间复杂度是 O(M * K)
*/

type TrieNode struct {
	Children [26]*TrieNode
	IsWord   bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children: [26]*TrieNode{},
		IsWord:   false,
	}
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) Insert(word string) {
	runes := []rune(word)
	node := t.root
	for _, r := range runes {
		idx := r - 'a'
		v := node.Children[idx]
		if v == nil {
			node.Children[idx] = NewTrieNode()
		}
		node = node.Children[idx]
	}
	node.IsWord = true
}

func (t *Trie) FindPrefix(word string) string {
	runes := []rune(word)
	node := t.root
	b := strings.Builder{}
	for _, r := range runes {
		idx := r - 'a'
		v := node.Children[idx]
		if v == nil || node.IsWord {
			break
		}
		b.WriteRune(r)
		node = node.Children[idx]
	}

	if node.IsWord {
		return b.String()
	} else {
		return ""
	}
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func replaceWords(dictionary []string, sentence string) string {
	trie := NewTrie()
	for _, d := range dictionary {
		trie.Insert(d)
	}

	b := strings.Builder{}
	words := strings.Split(sentence, " ")
	for _, word := range words {
		prefixWord := trie.FindPrefix(word)
		if prefixWord == "" {
			b.WriteString(word)
		} else {
			b.WriteString(prefixWord)
		}
		b.WriteString(" ")
	}
	res := b.String()
	res = strings.TrimRight(res, " ")

	return res
}
