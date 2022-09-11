package l21

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

func ListToArray(head *ListNode) []int {
	var res = make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

func NewList(vals []int) *ListNode {
	var nodeMap = make(map[int]*ListNode)

	for idx, val := range vals {
		node := NewNode(val, nil)
		nodeMap[idx] = node

		if idx > 0 {
			nodeMap[idx-1].Next = node
		}
	}

	return nodeMap[0]
}

/*
## 解题思路

循环遍历两个链表，依次将较小的节点放进来即可

## 复杂度分析

令 l1 的长度是 M，l2 的长度是 N

时间复杂度: O(M + N)
空间复杂度: O(1)，只申请了两个指向节点的指针，复杂度为常数级
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var (
		// res 是一个哨兵节点
		res = &ListNode{}
	)

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var resPos = res
	for list1 != nil && list2 != nil {
		if list2.Val < list1.Val {
			resPos.Next = list2
			list2 = list2.Next
		} else {
			resPos.Next = list1
			list1 = list1.Next
		}

		resPos = resPos.Next
	}

	if list1 != nil {
		resPos.Next = list1
	}
	if list2 != nil {
		resPos.Next = list2
	}

	return res.Next
}
