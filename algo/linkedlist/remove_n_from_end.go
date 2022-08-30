package linkedlist

var (
	cur = 0
)

/*removeNthFromEnd

## 题目说明

给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

## 解题思路

1. 把 current 之后的节点都执行一遍 删除倒数第N 的逻辑 (这样 cur 就执行过+1了，cur 就表示当前节点是倒数第几个节点)
2. 判断当前节点是不是倒数第 N 个节点，
	- 如果是，则删除，向上返回当前节点下一个
	- 如果不是，向上返回当前节点
	- 这样就为当前节点前一个执行了 删除倒数第N 的逻辑
	- 如此递归，直到递归到 head 节点，这样整个链表的每个节点都执行了 删除倒数第N 的逻辑

## 复杂度分析

N 表示链表的长度

时间复杂度 O(N)
空间复杂度 O(N)，每个节点都会调用一次 removeNthFromEnd 函数，创建了函数栈空间

这个解法的复杂度并没有太优秀，我用个 map[int]SingleListNode[T] 存储每个节点的指针，然后根据索引执行删除，时间和空间复杂度也都是 O(N)

但是它很帅，看起来比写 map 存索引的方式帅多了
*/
func removeNthFromEnd[T any](head *SingleListNode[T], n int) *SingleListNode[T] {
	if head == nil {
		return nil
	}

	head.Next = removeNthFromEnd(head.Next, n)

	cur++
	if cur == n {
		head = head.Next
	}

	return head
}
