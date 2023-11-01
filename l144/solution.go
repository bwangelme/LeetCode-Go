package l144

/*
## 题目 反转二叉树

给定一棵二叉树的根节点 root，请左右翻转这棵二叉树，并返回其根节点。

提示：
    树中节点数目范围在 [0, 100] 内
    -100 <= Node.val <= 100

## 复杂度分析

N 是树中节点的个数

### 递归解法

由于会遍历每个节点，所以时间复杂度是 O(N)

空间复杂度，最坏情况下树是一个链表，此时递归深度是 N, 空间复杂度是 O(N)

### 栈解法

因为会遍历每个节点，时间复杂度是 O(N)

空间复杂度，最坏情况下树是满二叉树，栈的大小是树的深度，
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// mirrorTreeRecursion
// 递归解法
func mirrorTreeRecursion(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmp := root.Left
	root.Left = mirrorTreeRecursion(root.Right)
	root.Right = mirrorTreeRecursion(tmp)

	return root
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		// stack pop
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.Left != nil {
			stack = append(stack, top.Left)
		}
		if top.Right != nil {
			stack = append(stack, top.Right)
		}

		top.Left, top.Right = top.Right, top.Left
	}

	return root
}
