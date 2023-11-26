package l210

/*
## 题目

现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。给你一个数组 prerequisites ，其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。

例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。

提示：
	1 <= numCourses <= 2000
	0 <= prerequisites.length <= numCourses * (numCourses - 1)
	prerequisites[i].length == 2
	0 <= ai, bi < numCourses
	ai != bi
	所有[ai, bi] 互不相同

## 解题思路

n 表示顶点的个数
v 表示边的个数

### 时间复杂度

O(v) + O(n) + O(n * v)

最坏情况下，针对每个顶点，都要执行一次遍历所有边并为它的指向顶点减少入度的操作

### 空间复杂度

O(n) + O(n)

用一个 map 存储所有顶点入度，O(n)
用一个 queue 最多存储所有顶点 O(n)

因此空间复杂度是 O(n)

## 复杂度分析
*/

type Queue struct {
	vals []int
	n    int
}

func NewQueue() *Queue {
	return &Queue{
		vals: make([]int, 0),
		n:    0,
	}
}

func (q *Queue) Enqueue(val int) {
	q.vals = append(q.vals, val)
	q.n++
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		return -1
	}

	val := q.vals[0]
	q.vals = q.vals[1:]
	q.n--
	return val

}

func (q *Queue) IsEmpty() bool {
	return q.n == 0
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		indegrees = make([]int, numCourses+1)
		res       = make([]int, numCourses)
		q         = NewQueue()
	)

	// 初始化所有顶点的入度
	for _, v := range prerequisites {
		indegrees[v[0]]++
	}

	// 将所有入度 == 0 的定点放入到队列中
	for i := 0; i < numCourses; i++ {
		if indegrees[i] == 0 {
			q.Enqueue(i)
		}
	}

	var i = 0
	for !q.IsEmpty() {
		curr := q.Dequeue()
		res[i] = curr
		i++

		// 注意边的方向是 v[1] --> v[0]
		for _, v := range prerequisites {
			if v[1] == curr {
				indegrees[v[0]]--
				if indegrees[v[0]] == 0 {
					q.Enqueue(v[0])
				}
			}
		}
	}

	if i == numCourses {
		return res
	} else {
		return make([]int, 0)
	}
}
