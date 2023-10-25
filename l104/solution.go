package l104

/*
## 题目

给定一个二叉树 root ，返回其最大深度。

二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。

## maxDepth 复杂度分析

令树的节点数为 N

时间复杂度是 O(N), 因为每个节点都会检查一次
空间复杂度是 O(N), 最坏情况下，树退化成链表，递归深度为树的节点长度 N

## maxDepthWidth 复杂度分析

令树的节点数为 N

时间复杂度是 O(N), 因为每个节点都会检查一次
空间复杂度是 O(N), 最坏情况下，树是满二叉树，queue 存储 叶子节点的个数, (n+1)/2
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepth
// 深度优先遍历
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// maxDepthWidth
// 广度优先遍历
func maxDepthWidth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		tmp := make([]*TreeNode, 0)
		depth += 1
		for _, node := range queue {
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		queue = tmp
	}

	return depth
}
