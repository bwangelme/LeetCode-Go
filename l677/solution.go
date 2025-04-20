package l677

/*
## 题目描述

设计一个 map ，满足以下几点:

字符串表示键，整数表示值
返回具有前缀等于给定字符串的键的值的总和
实现一个 MapSum 类：

MapSum() 初始化 MapSum 对象
void insert(String key, int val) 插入 key-val 键值对，字符串表示键 key ，整数表示值 val 。如果键 key 已经存在，那么原来的键值对 key-value 将被替代成新的键值对。
int sum(string prefix) 返回所有以该前缀 prefix 开头的键 key 的值的总和。


示例 1：

输入：
["MapSum", "insert", "sum", "insert", "sum"]
[[], ["apple", 3], ["ap"], ["app", 2], ["ap"]]
输出：
[null, null, 3, null, 5]

解释：
MapSum mapSum = new MapSum();
mapSum.insert("apple", 3);
mapSum.sum("ap");           // 返回 3 (apple = 3)
mapSum.insert("app", 2);
mapSum.sum("ap");           // 返回 5 (apple + app = 3 + 2 = 5)


提示：

1 <= key.length, prefix.length <= 50
key 和 prefix 仅由小写英文字母组成
1 <= val <= 1000
最多调用 50 次 insert 和 sum

## 解题思路

1. 将所有单词构造成一个前缀树，每个节点中有一个 val 表示单词的权重，如果当前节点不是叶子节点，val = 0
2. 首先根据前缀遍历前缀树，找到前缀的下一个节点
3. 从前缀的下一个节点开始，遍历到叶子节点的所有和，将和加起来。遍历的方式是 dfs

## 复杂度分析

N 单词的个数
M 所有单词的平均长度
K 前缀的长度

### 空间复杂度

空间复杂度即为前缀树的空间复杂度，最坏复杂度是 O(N * M)

### 时间复杂度

遍历前缀所需的时间为 O(K)
从前缀开始，遍历所有叶子节点，时间复杂度为 O(N * (M-K))

因此时间复杂度为 O(N * (M-K))
*/

type TrieNode struct {
	Children [26]*TrieNode
	Val      int
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children: [26]*TrieNode{},
		Val:      0,
	}
}

type MapSum struct {
	Root *TrieNode
}

func Constructor() MapSum {
	return MapSum{
		Root: NewTrieNode(),
	}
}

func (this *MapSum) Insert(key string, val int) {
	node := this.Root
	chars := []byte(key)
	for _, c := range chars {
		idx := c - 'a'
		if node.Children[idx] == nil {
			node.Children[idx] = NewTrieNode()
		}
		node = node.Children[idx]
	}

	node.Val = val
}

func (this *MapSum) Sum(prefix string) int {
	node := this.Root
	chars := []byte(prefix)
	for _, c := range chars {
		idx := c - 'a'
		if node.Children[idx] == nil {
			return 0
		}
		node = node.Children[idx]
	}
	return this.getSum(node)
}

func (this *MapSum) getSum(node *TrieNode) int {
	if node == nil {
		return 0
	}

	var result = node.Val
	for _, c := range node.Children {
		result += this.getSum(c)
	}
	return result
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
