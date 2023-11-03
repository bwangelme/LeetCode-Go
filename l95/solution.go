package l95

/*
## 题目

给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。


    1 <= n <= 8

## 复杂度分析

时间复杂度:
*/

import (
	"github.com/bwangelme/LeetCode-Go/lt"
)

func generateTrees(n int) []*lt.TreeNode {
	return genTree(1, n)
}

// TODO 剪枝，将一些重复的结果缓存起来
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
		leftTrees := genTree(start, root-1)
		rightTrees := genTree(root+1, end)

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

func generateTreesCache(n int) []*lt.TreeNode {
	if n == 1 {
		return []*lt.TreeNode{&lt.TreeNode{Val: 1}}
	}
	cache := make([][][]*lt.TreeNode, n+1) // cache[i][j]表示数字i到j生成的字数
	for i := range cache {
		cache[i] = make([][]*lt.TreeNode, n+1)
		cache[i][i] = append(cache[i][i], &lt.TreeNode{Val: i})
	}
	var fn func(t, left, right int)
	fn = func(t, left, right int) {
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
		if t == left {
			for i := range cache[t+1][right] {
				cache[left][right] = append(cache[left][right], &lt.TreeNode{Val: t, Right: cache[t+1][right][i]})
			}
		} else if t == right {
			for i := range cache[left][t-1] {
				cache[left][right] = append(cache[left][right], &lt.TreeNode{Val: t, Left: cache[left][t-1][i]})
			}
		} else {
			for i := range cache[left][t-1] {
				for j := range cache[t+1][right] {
					cache[left][right] = append(cache[left][right], &lt.TreeNode{Val: t, Left: cache[left][t-1][i], Right: cache[t+1][right][j]})
				}
			}
		}
	}
	for i := 1; i <= n; i++ {
		fn(i, 1, n)
	}
	return cache[1][n]
}
