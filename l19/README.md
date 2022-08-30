删除链表的倒数第 N 个结点
===

见 [algo/linkedlist/remove_n_from_end.go](../algo/linkedlist/remove_n_from_end.go)

不知道 leetcode 的 OJ 是如何实现的，如果定义了全局变量，程序就会出错，于是我把 current 改成了函数的参数，这样结果就能通过了

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNFromEnd(head *ListNode, n int, current *int) *ListNode {
    if head == nil {
        return nil
    }

    head.Next = removeNFromEnd(head.Next, n, current)

    *current = *current + 1
    if *current == n {
        head = head.Next
    }
    return head
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var current int = 0
    return removeNFromEnd(head, n, &current)
}
```