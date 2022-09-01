package l141_l142

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l ListNode) String() string {
	return fmt.Sprintf("%v", l.Val)
}

func NewNode(val int, next *ListNode) *ListNode {
	return &ListNode{
		Val:  val,
		Next: next,
	}
}

func NewList(vals []int, pos int) *ListNode {
	var nodeMap = make(map[int]*ListNode)

	for idx, val := range vals {
		node := NewNode(val, nil)
		nodeMap[idx] = node

		if idx > 0 {
			nodeMap[idx-1].Next = node
		}
	}

	if pos >= 0 {
		nodeMap[len(vals)-1].Next = nodeMap[pos]
	}

	return nodeMap[0]
}

/*
## 解题思路

使用一个 map 存储每个节点的指针，遍历链表，当遍历到已经存在过的节点时，即证明此链表有环。

## 复杂度分析

这里的 N 都是代表链表的长度

时间复杂度 O(N)，会循环 N+1 次，遇到第一个重复的节点退出
空间复杂度 O(N)，申请了一个 map，map 的大小为链表的长度

*/
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	//iteredIdx 表示遍历过的索引
	var iteredIdx = make(map[*ListNode]struct{})

	for head != nil {
		_, exist := iteredIdx[head]
		if exist {
			return true
		}
		iteredIdx[head] = struct{}{}

		head = head.Next
	}

	return false
}

/*
## 说明

判断链表是否有环的空间复杂度为 O(1) 的解法

## 解题思路

创建快慢两个指针，慢指针每次走一步，快指针每次走两步

快慢指针进入环的位置不同，假设它们在环中的位置差为 M(始终认为慢在前，快在后) (0 <= M < L, L 为环的长度)

每次快比慢多走一步，所以走了 M 步之后，快就能追上慢

## 复杂度分析

只申请了两个指针，空间复杂度为 O(1)

时间复杂度为 O(N)，N 为链表的长度

环中循环执行的步骤取决与快慢在环中的位置差M，(0 <= M < L, L 为环的长度)
	最好为 M == 1，循环 1 次
	最坏为 M == 0，循环 L 次

简化可得时间复杂度为 O(N)
*/
func hasCycleO1(head *ListNode) bool {
	if head == nil {
		return false
	}

	var (
		fast = head
		slow = head
	)

	for fast != nil && fast.Next != nil && slow != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}

func detectCycle(head *ListNode) *ListNode {
	// TODO
	if head == nil {
		return nil
	}

	var (
		fast = head
		slow = head
	)

	for fast != nil && fast.Next != nil && slow != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return slow
		}
	}

	return nil
}
