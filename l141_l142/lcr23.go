package l141_l142

/*
## 题目

给定两个单链表的头节点 headA 和 headB ，请找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。

题目数据 保证 整个链式结构中不存在环。

注意，函数返回结果后，链表必须 保持其原始结构 。

提示：

- listA 中节点数目为 m
- listB 中节点数目为 n
- 0 <= m, n <= 3 * 10^4
- 1 <= Node.val <= 10^5
- 0 <= skipA <= m
- 0 <= skipB <= n
- 如果 listA 和 listB 没有交点，intersectVal 为 0
- 如果 listA 和 listB 有交点，intersectVal == listA[skipA + 1] == listB[skipB + 1]

## 解题思路

1. 求出 a, b 两个链表的长度, 和长度的差值 delta
2. p1 在较长的链表上移动 delta 步
3. p2 指向较短链表的 head, 然后 p1 和 p2 一起开始移动
4. 如果 p1 和 p2 相等时，就是两个链表的相交点
5. 如果两个链表没有交点，则 p1 和 p2 也会遍历到链表的尾部，然后 p1 == p2 break

## 复杂度分析

时间复杂度:

O(m + n + delta + max(m, n)) = O(m + n)

空间复杂度:

O(1)
*/

import (
	"math"
)

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var (
		cntA     = 0
		cntB     = 0
		currentA = headA
		currentB = headB
	)

	for currentA != nil {
		currentA = currentA.Next
		cntA += 1
	}
	for currentB != nil {
		currentB = currentB.Next
		cntB += 1
	}

	var (
		delta      = int(math.Abs(float64(cntA - cntB)))
		firstMove  = headA
		secondMove = headB
	)
	if cntA < cntB {
		firstMove = headB
		secondMove = headA
	}

	for delta > 0 {
		firstMove = firstMove.Next
		delta--
	}

	for {
		if firstMove == secondMove {
			break
		}
		firstMove = firstMove.Next
		secondMove = secondMove.Next
	}

	return firstMove
}
