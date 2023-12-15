package linkedlist

/*
本程序演示了哨兵节点如何简化链表节点的编写
*/

func Append[T any](head *SingleListNode[T], val T) *SingleListNode[T] {
	var node = &SingleListNode[T]{
		Next: nil,
		Data: val,
	}

	if head == nil {
		return node
	}

	var current = head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node

	return head
}

func AppendDummy[T any](head *SingleListNode[T], val T) *SingleListNode[T] {
	var node = &SingleListNode[T]{
		Next: nil,
		Data: val,
	}
	var dummy = &SingleListNode[T]{
		Next: head,
	}
	var current = dummy

	for current.Next != nil {
		current = current.Next
	}

	current.Next = node

	return dummy.Next
}

func DeleteVal[T string](head *SingleListNode[T], val T) *SingleListNode[T] {
	if head == nil {
		return head
	}

	if head.Data == val {
		return head.Next
	}

	var current = head
	for current.Next != nil {
		if current.Next.Data == val {
			current.Next = current.Next.Next
			break
		}
		current = current.Next
	}

	return head
}

func DeleteValDummy[T string](head *SingleListNode[T], val T) *SingleListNode[T] {
	var dummy = &SingleListNode[T]{
		Next: head,
	}

	var current = dummy
	for current.Next != nil {
		if current.Next.Data == val {
			current.Next = current.Next.Next
			break
		}
		current = current.Next
	}

	return dummy.Next
}
