package lcr029

type Node struct {
	Val  int
	Next *Node
}

func insert(head *Node, val int) *Node {
	node := &Node{
		Val:  val,
		Next: nil,
	}
	if head == nil {
		head = node
		node.Next = node
	} else if head.Next == head {
		head.Next = node
		node.Next = head
	} else {
		insertCore(head, node)
	}

	return head
}

func insertCore(head *Node, node *Node) {
	var (
		cur     = head
		next    = head.Next
		biggest = head
	)
	for !(cur.Val <= node.Val && node.Val < next.Val) && next != head {
		cur = next
		next = next.Next
		if cur.Val >= biggest.Val {
			biggest = cur
		}
	}

	if cur.Val <= node.Val && node.Val <= next.Val {
		node.Next = cur.Next
		cur.Next = node
	} else {
		node.Next = biggest.Next
		biggest.Next = node
	}
}
