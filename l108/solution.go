package l108

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
## 题目

给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。

高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。

提示：

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums 按 严格递增 顺序排列

## 复杂度分析

# N 表示数组的长度

### 时间复杂度

sortedArrayToBST 函数的调用是一颗二叉树

例如数组长度是 7

	s(7)

s(3)            s(3)
s(1) s(1)       s(1) s(1)

树的总节点数 N， 所以 函数的调用次数也是 N， 时间复杂度是 O(N)

### 空间复杂度

函数调用完成后，会释放空间，因此空间复杂度主要考虑树的高度，因此空间复杂度是 O(logN)
*/
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	middle := len(nums) / 2
	root := nums[middle]
	left := sortedArrayToBST(nums[0:middle])
	right := sortedArrayToBST(nums[middle+1:])
	rootNode := &TreeNode{
		Val:   root,
		Left:  left,
		Right: right,
	}

	return rootNode
}
