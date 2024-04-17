package l297

import (
	"fmt"
	"github.com/bwangelme/LeetCode-Go/lt"
	"strconv"
	"strings"
)

/*
## 题目描述

297. 二叉树的序列化与反序列化

序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

提示：

树中结点数在范围 [0, 104] 内
-1000 <= Node.val <= 1000

## 解题思路

利用先序遍历的方式将树中所有节点输出，空节点输出为 #

## 复杂度分析

n == 树中节点个数

serialize

时间复杂度 O(N)

空间复杂度

递归深度为树的高度，最好是 O(logN), 最坏是 O(N)

deserialize

时间复杂度 O(N)

空间复杂度 O(N), nodeStrs 占用的空间大小是 O(N)

*/

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *lt.TreeNode) string {
	if root == nil {
		return "#"
	}

	leftStr := this.serialize(root.Left)
	rightStr := this.serialize(root.Right)
	return fmt.Sprintf("%d,%s,%s", root.Val, leftStr, rightStr)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *lt.TreeNode {
	nodeStrs := strings.Split(data, ",")
	var i = 0
	return dfs(nodeStrs, &i)
}

func dfs(nodeStrs []string, i *int) *lt.TreeNode {
	v := nodeStrs[*i]
	*i++

	if v == "#" {
		return nil
	}

	vNum, _ := strconv.Atoi(v)
	node := &lt.TreeNode{
		Val: vNum,
	}
	node.Left = dfs(nodeStrs, i)
	node.Right = dfs(nodeStrs, i)
	return node
}
