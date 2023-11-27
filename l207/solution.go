package l207

/*
## 题目描述

你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。

在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。

    例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。

请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

提示：

    1 <= numCourses <= 2000
    0 <= prerequisites.length <= 5000
    prerequisites[i].length == 2
    0 <= ai, bi < numCourses
    prerequisites[i] 中的所有课程对 互不相同

## 解题思路

和 210 是完全相同的解题思路

## 复杂度分析

v 图中节点的个数
e 图中的边的个数

### 时间复杂度

O(e) + O(v) + O(v * e) = O(v * e)

最坏情况下，对于每个节点，都会执行

for _, v := range prerequisites

寻找当前节点的连接的节点，将其入度-1,并判断是否可以加入到队列中

### 空间复杂度

indegrees 中存储了 v 个节点的入度
q 中最多存储了 v 个节点

因此空间复杂度是 O(v)

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

func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		indegrees = make([]int, numCourses+1)
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

	return i == numCourses
}
