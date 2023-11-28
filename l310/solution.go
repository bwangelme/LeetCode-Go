package l310

/*
## 题目

树是一个无向图，其中任何两个顶点只通过一条路径连接。 换句话说，一个任何没有简单环路的连通图都是一棵树。

给你一棵包含 n 个节点的树，标记为 0 到 n - 1 。给定数字 n 和一个有 n - 1 条无向边的 edges 列表（每一个边都是一对标签），其中 edges[i] = [ai, bi] 表示树中节点 ai 和 bi 之间存在一条无向边。

可选择树中任何一个节点作为根。当选择节点 x 作为根节点时，设结果树的高度为 h 。在所有可能的树中，具有最小高度的树（即，min(h)）被称为 最小高度树 。

请你找到所有的 最小高度树 并按 任意顺序 返回它们的根节点标签列表。
树的 高度 是指根节点和叶子节点之间最长向下路径上边的数量。

提示：

    1 <= n <= 2 * 10^4
    edges.length == n - 1
    0 <= ai, bi < n
    ai != bi
    所有 (ai, bi) 互不相同
    给定的输入 保证 是一棵树，并且 不会有重复的边

## 解题思路

这题用广度优先遍历的思路来解，从出度 == 1 的节点开始，一层一层地删除节点，最后剩下的一层，就是能构成最小高度树的叶子节点

越靠近里面的节点，越有可能是最小高度树

1. 不断去掉所有出度为1的节点
2. 将相邻节点的出度 -1

最后一批出度为1的节点, 就是能构成最小高度树的节点

## 复杂度分析

v 表示节点个数
e 表示边的个数

### 时间复杂度

// 初始化操作
O(e) + O(v)

// 所有节点都会放入到队列中,
O(v)
所有出度为1的节点的相邻节点都会只有1个, 所以每次遍历的邻居只有1个

总的时间复杂度是 O(e) + O(v)
*/

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	var (
		degrees   = make(map[int]int)
		neighbors = make(map[int][]int)
		queue     []int
	)

	// 计算每个节点的出度和邻居
	for _, edge := range edges {
		degrees[edge[0]] += 1
		degrees[edge[1]] += 1

		neighbors[edge[0]] = append(neighbors[edge[0]], edge[1])
		neighbors[edge[1]] = append(neighbors[edge[1]], edge[0])
	}

	// 将出度为1的节点放入到队列中
	for node, degree := range degrees {
		if degree == 1 {
			queue = append(queue, node)
		}
	}

	var res []int
	for len(queue) > 0 {
		// res 每次存储的，都是每一层的叶子节点
		res = make([]int, 0)
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			res = append(res, cur)

			// 这里会遍历一些已经删除的节点, 将它们的出度减少成0,不影响结果
			for _, neighbor := range neighbors[cur] {
				degrees[neighbor]--
				if degrees[neighbor] == 1 {
					queue = append(queue, neighbor)
				}
			}
		}
		// 将刚刚遍历过的节点，全部删掉
		queue = queue[size:]
	}

	// 将最后一层的叶子节点返回，这也就是所有最后的出度 == 1 的节点，它们可以构成最小高度树
	return res
}
