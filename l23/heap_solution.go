package l23

import (
	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/emirpasic/gods/utils"
)

/*
## 题目描述

合并 K 个升序链表

## 解题思路

将链表的头节点加入到最小堆中，每次从堆中取出最小值，如果最小值还有next,将 next 加入到堆中

堆中所有节点遍历完成后，最小堆也会变成空

## 复杂度分析

K = 链表的数量
N = 链表的总长度

- 空间复杂度 O(K), 最多同时存储 k 个节点
- 时间复杂度

最小堆每次取节点的时间复杂度，O(1)
最小堆每次添加节点的时间复杂度 O(logK)

因此一共执行 N 次添加节点和取节点，总的时间复杂度是 O(NlogK)
*/
func mergeKListsByHeap(lists []*ListNode) *ListNode {
	var (
		dummy = &ListNode{}
		cur   = dummy
	)
	var (
		minHeap = binaryheap.NewWith(func(a, b interface{}) int {
			aNode := a.(*ListNode)
			bNode := b.(*ListNode)
			return utils.IntComparator(aNode.Val, bNode.Val)
		})
	)

	for _, list := range lists {
		if list != nil {
			minHeap.Push(list)
		}
	}

	for !minHeap.Empty() {
		v, _ := minHeap.Pop()
		node := v.(*ListNode)

		cur.Next = node
		cur = node

		if node.Next != nil {
			minHeap.Push(node.Next)
		}
	}

	return dummy.Next
}
