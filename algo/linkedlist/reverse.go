package linkedlist

/*
Reverse 反转单链表

## 注意事项

1. 注意 head.next 要被设置成 nil
2. 循环结束后，head 是 nil，prev 表示原来链表的尾部

## 复杂度分析

时间复杂度是 O(N), N 表示链表的长度
*/
func Reverse[T any](head *SingleListNode[T]) (tail *SingleListNode[T]) {
	var prev *SingleListNode[T] = nil

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}
