package l102

var res = make([][]int, 0)

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recursiveLevelOrder(node *TreeNode, level int) {
	if node == nil {
		return
	}

	if len(res) < level {
		item := make([]int, 0)
		res = append(res, item)
	}

	item := res[level-1]
	item = append(item, node.Val)
	res[level-1] = item

	recursiveLevelOrder(node.Left, level+1)
	recursiveLevelOrder(node.Right, level+1)
}

// 时间复杂度 & 空间复杂度
// 计算本程序的复杂度，主要看 recursiveLevelOrder 的执行次数。
// 因为 recursiveLevelOrder 中没有 for 循环，所以 recursiveLevelOrder 的复杂度可以看做是1
// 函数 recursiveLevelOrder 运行的次数等于二叉树的深度，设树的节点为N
// 最坏情况下树的深度为 N
// 最好情况下(满二叉树)树的深度为 log2(N+1)
func levelOrder(root *TreeNode) [][]int {
	recursiveLevelOrder(root, 1)
	return res
}
