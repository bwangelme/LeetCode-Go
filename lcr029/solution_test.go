package lcr029

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func (n *Node) String() string {
	var buf = new(bytes.Buffer)
	head := n
	cur := head.Next

	fmt.Fprintf(buf, "%v", head.Val)
	for cur != head {
		fmt.Fprintf(buf, " -> %v", cur.Val)
		cur = cur.Next
	}

	return buf.String()
}

func createLoopList(nums []int) *Node {
	head := &Node{
		Val:  nums[0],
		Next: nil,
	}
	prevNode := head

	for _, val := range nums[1:] {
		node := &Node{
			Val:  val,
			Next: nil,
		}
		prevNode.Next = node
		prevNode = node
	}

	prevNode.Next = head

	return head
}

func Test_Insert(t *testing.T) {
	for _, tt := range []struct {
		Arr    []int
		Val    int
		Wanted string
	}{
		{
			[]int{3, 4, 1},
			2,
			"3 -> 4 -> 1 -> 2",
		},
		{
			[]int{1, 3, 3},
			4,
			"1 -> 3 -> 3 -> 4",
		},
		{
			[]int{1, 3, 3, 4},
			3,
			"1 -> 3 -> 3 -> 3 -> 4",
		},
	} {
		head := createLoopList(tt.Arr)
		res := insert(head, tt.Val)
		assert.Equal(t, tt.Wanted, fmt.Sprint(res))
	}
}
