package l95

/*
## 题目

给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。

    1 <= n <= 8

## genTree 复杂度分析

时间复杂度：
	整个算法的时间复杂度取决于「可行二叉搜索树的个数」，而对于 n 个点生成的二叉搜索树数量等价于数学上第 n 个「卡特兰数」，
用 Gn 表示。卡特兰数具体的细节请读者自行查询，这里不再赘述，只给出结论。
生成一棵二叉搜索树需要 O(n) 的时间复杂度，一共有 Gn 棵二叉搜索树，也就是 O(nGn)。
而卡特兰数以 4^n/n^{3/2} 增长，因此总时间复杂度为 O(4^n/n^{1/2})。

空间复杂度：
	n 个点生成的二叉搜索树有 Gn 棵，每棵有 n 个节点，
因此存储的空间需要 O(nGn) = O(4^n/n^{1/2}) ，递归函数需要 O(n) 的栈空间，
因此总空间复杂度为 O(4^n/n^{1/2})
*/

import (
	"github.com/bwangelme/LeetCode-Go/lt"
)

func generateTrees(n int) []*lt.TreeNode {
	return genTree(1, n)
}

func genTree(start, end int) []*lt.TreeNode {
	var res = make([]*lt.TreeNode, 0)
	// 这一句让, leftTrees 和 rightTrees 返回的长度不是0, 而是 长度为1, 里面的值是 nil
	if start > end {
		return []*lt.TreeNode{nil}
	}

	for root := start; root <= end; root++ {
		leftTrees := genTree(start, root-1)
		rightTrees := genTree(root+1, end)

		for _, left := range leftTrees {
			for _, right := range rightTrees {
				tree := &lt.TreeNode{
					Val:   root,
					Left:  left,
					Right: right,
				}
				res = append(res, tree)
			}
		}

	}

	return res
}

// genTreeTedious 和 genTree 是同样的思路，但是因为没有哨兵，判断起来很麻烦
func genTreeTedious(start, end int) []*lt.TreeNode {
	var res = make([]*lt.TreeNode, 0)

	for root := start; root <= end; root++ {
		leftTrees := genTreeTedious(start, root-1)
		rightTrees := genTreeTedious(root+1, end)

		if len(leftTrees) == 0 && len(rightTrees) == 0 {
			tree := &lt.TreeNode{
				Val:   root,
				Left:  nil,
				Right: nil,
			}
			res = append(res, tree)
		} else if len(leftTrees) == 0 {
			for _, right := range rightTrees {
				tree := &lt.TreeNode{
					Val:   root,
					Left:  nil,
					Right: right,
				}
				res = append(res, tree)
			}
		} else if len(rightTrees) == 0 {
			for _, left := range leftTrees {
				tree := &lt.TreeNode{
					Val:   root,
					Left:  left,
					Right: nil,
				}
				res = append(res, tree)
			}
		} else {
			for _, left := range leftTrees {
				for _, right := range rightTrees {
					tree := &lt.TreeNode{
						Val:   root,
						Left:  left,
						Right: right,
					}
					res = append(res, tree)
				}
			}

		}

	}

	return res
}

func genTreeCache(n int) []*lt.TreeNode {
	// n = 1 时只有一颗树
	if n == 1 {
		return []*lt.TreeNode{{Val: n}}
	}

	// cache[i][j]表示数字i到j生成的子树
	var cache = make([][][]*lt.TreeNode, n+1)

	// 填满叶子节点
	for i := 1; i <= n; i++ {
		cache[i] = make([][]*lt.TreeNode, n+1)
		cache[i][i] = []*lt.TreeNode{{Val: i}}
	}

	// fn 表示以 t 为根节点，算出 [left,right] 中的所有子树
	var fn func(t, left, right int)
	fn = func(t, left, right int) {
		// 先填满左子树和右子树
		if t != left && len(cache[left][t-1]) == 0 {
			for i := left; i <= t-1; i++ {
				fn(i, left, t-1)
			}
		}
		if t != right && len(cache[t+1][right]) == 0 {
			for i := t + 1; i <= right; i++ {
				fn(i, t+1, right)
			}
		}

		// 再补充满 cache[left][right]
		if t == left {
			for i := range cache[t+1][right] {
				cache[left][right] = append(cache[left][right], &lt.TreeNode{
					Val:   t,
					Right: cache[t+1][right][i],
				})
			}
		} else if t == right {
			for i := range cache[left][t-1] {
				cache[left][right] = append(cache[left][right], &lt.TreeNode{
					Val:  t,
					Left: cache[left][t-1][i],
				})
			}
		} else {
			for i := range cache[left][t-1] {
				for j := range cache[t+1][right] {
					cache[left][right] = append(cache[left][right], &lt.TreeNode{
						Val:   t,
						Left:  cache[left][t-1][i],
						Right: cache[t+1][right][j],
					})
				}
			}
		}
	}

	// 分别以 1-n 中的数为根节点，算出 1-n 的所有子树
	for i := 1; i <= n; i++ {
		fn(i, 1, n)
	}

	// 返回 1-n 的所有子树
	return cache[1][n]
}
