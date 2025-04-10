package lt

/*
## 描述

实现支持 Unicode 字符的前缀树

*/

type UTrieNode struct {
	Children map[rune]*UTrieNode
	IsWord   bool
}

func NewUTrieNode() *UTrieNode {
	return &UTrieNode{
		IsWord:   false,
		Children: make(map[rune]*UTrieNode, 0),
	}
}

type UTrie struct {
	root *UTrieNode
}

func NewUTrie() UTrie {
	return UTrie{
		root: NewUTrieNode(),
	}
}

func (u *UTrie) Insert(word string) {
	runes := []rune(word)
	node := u.root

	for _, r := range runes {
		_, ok := node.Children[r]
		if !ok {
			node.Children[r] = NewUTrieNode()
		}

		node = node.Children[r]
	}
	node.IsWord = true
}

func (u *UTrie) Search(word string) bool {
	runes := []rune(word)

	node := u.root
	for _, r := range runes {
		v, ok := node.Children[r]
		if !ok {
			return false
		}
		node = v
	}

	return node.IsWord
}

func (u *UTrie) StartsWith(prefix string) bool {
	runes := []rune(prefix)

	node := u.root
	for _, r := range runes {
		v, ok := node.Children[r]
		if !ok {
			return false
		}
		node = v
	}

	return true
}
