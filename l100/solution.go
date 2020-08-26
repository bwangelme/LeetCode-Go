package l100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

/**
 * 设较少节点的树的节点数是N，最坏情况下，时间复杂度是 O(N+1)
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		if p == q {
			return true
		}
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if !isSameTree(p.Left, q.Left) {
		return false
	}

	if !isSameTree(p.Right, q.Right) {
		return false
	}

	return true
}
