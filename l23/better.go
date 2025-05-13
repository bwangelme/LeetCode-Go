package l23

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{} // 用哨兵节点简化代码逻辑
	cur := dummy         // cur 指向新链表的末尾
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1 // 把 list1 加到新链表中
			list1 = list1.Next
		} else { // 注：相等的情况加哪个节点都是可以的
			cur.Next = list2 // 把 list2 加到新链表中
			list2 = list2.Next
		}
		cur = cur.Next
	}
	// 拼接剩余链表
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return dummy.Next
}

/*
## 解题思路

使用归并排序的思路，将链表二分地进行排序，然后将两个排序后的链表进行合并

## 复杂度分析

K = 链表的数量
N = 节点的总数量

栈的递归深度是 O(logK),因此空间复杂度是 O(logK)

- 时间复杂度

一共执行了 O(logK) 层的 mergeTwoLists 函数
每层执行的时候，都是要遍历 N 个节点，因此时间复杂度是 O(NlogK)
*/
func mergeKListsBetter(lists []*ListNode) *ListNode {
	m := len(lists)
	if m == 0 { // 注意输入的 lists 可能是空的
		return nil
	}
	if m == 1 { // 无需合并，直接返回
		return lists[0]
	}
	left := mergeKListsBetter(lists[:m/2])  // 合并左半部分
	right := mergeKListsBetter(lists[m/2:]) // 合并右半部分
	return mergeTwoLists(left, right)       // 最后把左半和右半合并
}
