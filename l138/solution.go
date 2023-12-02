package l138

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func printList(head *Node) {
	fmt.Println("--------------")
	for head != nil {
		fmt.Printf("val: %v", head.Val)
		if head.Random != nil {
			fmt.Printf(" random: %v\n", head.Random.Val)
		} else {
			fmt.Printf(" random: nil\n")
		}
		head = head.Next
	}
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	var current = head
	for current != nil {
		copyNode := &Node{
			Val:    current.Val,
			Next:   current.Next,
			Random: current.Random,
		}
		current.Next = copyNode
		current = copyNode.Next
	}

	current = head
	for current != nil {
		copyNode := current.Next
		if current.Random != nil {
			copyNode.Random = current.Random.Next
		}
		current = copyNode.Next
	}

	current = head
	res := head.Next
	for current != nil {
		copyNode := current.Next
		current.Next = copyNode.Next

		if copyNode.Next != nil {
			copyNode.Next = copyNode.Next.Next
		}

		current = current.Next
	}

	return res
}
