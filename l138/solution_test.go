package l138

import (
	"testing"
)

func makeList(nodes [][]int) *Node {
	var head = &Node{}
	var nodesMap = make(map[int]*Node)

	var current = head
	for idx, node := range nodes {
		tmp := &Node{
			Val:    node[0],
			Next:   nil,
			Random: nil,
		}
		nodesMap[idx+1] = tmp
		current.Next = tmp
		current = current.Next
	}

	current = head.Next
	for _, node := range nodes {
		current.Random = nodesMap[node[1]]
		current = current.Next
	}

	return head.Next
}

func Test_copyRandomList(t *testing.T) {
	head := makeList([][]int{{1, 1}, {2, 1}})
	printList(head)
	res := copyRandomList(head)
	printList(res)
}
