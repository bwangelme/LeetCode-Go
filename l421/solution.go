package l421

/*
## 题目描述

给你一个整数数组 nums ，返回 nums[i] XOR nums[j] 的最大运算结果，其中 0 ≤ i ≤ j < n 。

示例 1：

输入：nums = [3,10,5,25,2,8]
输出：28
解释：最大运算结果是 5 XOR 25 = 28.
示例 2：

输入：nums = [14,70,53,83,49,91,36,80,92,51,66,70]
输出：127


提示：

1 <= nums.length <= 2 * 105
0 <= nums[i] <= 231 - 1

## 解题思路

1. 将所有数字，从高位到低位，分解成 32 个 0,1 位数组
2. 将这些数字的位数组组合成一个前缀树，每个前缀树节点只有两个子节点 0 或 1
3. 遍历每个数字，
	在前缀树中寻找每一位不同的数字，如果存在，则这一位的异或结果是1,否则是0
	当前数字遍历的结果就是，包含当前数字的异或结果最大值
4. 遍历完成后，将每个数字的异或结果最大值进行比较，得到数组的任意两个数字异或结果最大值

## 复杂度分析

N = 数字的个数

### 空间复杂度

空间复杂度即为前缀树的空间，O(32 * N) = O(N)

### 时间复杂度

构造前缀树 buildTrie,针对每个数字遍历32位，时间复杂度是 O(32 * N) = O(N)
查找异或最大值 findMaximumXOR, 针对每个数字遍历前缀树，查找异或最大值，时间复杂度是 O(32 * N) = O(N)

因此总的时间复杂度是 O(N)
*/

type TrieNode struct {
	Children [2]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children: [2]*TrieNode{},
	}
}

func buildTrie(nums []int) *TrieNode {
	root := NewTrieNode()
	for _, num := range nums {
		node := root
		for i := 31; i >= 0; i-- {
			bit := (num >> i) & 1
			if node.Children[bit] == nil {
				node.Children[bit] = NewTrieNode()
			}
			node = node.Children[bit]
		}
	}

	return root
}

func findMaximumXOR(nums []int) int {
	root := buildTrie(nums)
	var max int = 0
	for _, num := range nums {
		node := root
		xor := 0
		for i := 31; i >= 0; i-- {
			bit := (num >> i) & 1
			// 执行异或运算的时候，如果两个数字的同一位不同，那么异或的结果更大，因此优先找同一位不同的数字 node.Children[1-bit] 表示和当前数字当前位不同的数字。
			//		找到的话，当前位的异或结果就是 1, xor = (xor << 1) + 1
			//		如果没找到的话，当前位的异或结果就是 0, xor = (xor << 1)
			if node.Children[1-bit] != nil {
				xor = (xor << 1) + 1
				node = node.Children[1-bit]
			} else {
				xor = xor << 1
				node = node.Children[bit]
			}
		}
		// xor 表示包含当前数字 num 的，两个数的异或结果最大值
		max = Max(max, xor)
	}

	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
