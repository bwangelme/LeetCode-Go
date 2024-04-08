package l94

/*
给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

输入：root = [1,null,2,3]
输出：[1,3,2]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]


提示：

树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100

进阶: 递归算法很简单，你可以通过迭代算法完成吗？

*/

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

/*
## 思路

1. 将树的左子树不断入栈
2. 从栈顶弹出的节点没有左子树或者左子树已经遍历过
3. 输出栈顶节点
4. 将栈顶节点的右子树执行上述 1 2 3 逻辑

这样实现了 左 根 右 的遍历逻辑

## 复杂度分析

- 空间复杂度

栈的大小不超过树的深度，所以空间复杂度是 O(logN), 最坏情况下，树是链表，此时复杂度是 O(N)

- 时间复杂度

最坏情况下，树是一条链表，此时时间复杂度是 O(2N) (入栈和出栈各算一次)，等于 O(N)
*/
func inorderTraversal(root *TreeNode) []int {
	var res = make([]int, 0)
	var stack = make([]*TreeNode, 0)

	for root != nil || len(stack) > 0 {
		// 将 root 所有的左叶子压入栈中
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		// pop stack
		stack = stack[0 : len(stack)-1]
		// 先记录左叶子
		res = append(res, root.Val)
		// 再去遍历右子树
		root = root.Right
	}

	return res
}

func inorderTraversal_recursion(root *TreeNode) []int {
	var res = make([]int, 0)

	if root == nil {
		return res
	}

	res = append(res, inorderTraversal_recursion(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal_recursion(root.Right)...)

	return res
}
