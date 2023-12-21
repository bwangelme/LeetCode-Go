package lcr28

/*
## 题目

多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。

给定位于列表第一级的头节点，请扁平化列表，即将这样的多级双向链表展平成普通的双向链表，使所有结点出现在单级双链表中。

提示：

节点数目不超过 1000
1 <= Node.val <= 10^5

## 复杂度分析

N 表示节点的个数
K 表示层级的深度

时间复杂度: 遍历了所有节点一遍，时间复杂度是 O(N)
空间复杂度：flattenGetTail 的调用深度和链表层级的深度一样，因此空间复杂度是 O(K)
*/

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	flattenGetTail(root)
	return root
}

func flattenGetTail(head *Node) *Node {
	var (
		node       = head
		tail *Node = nil
	)

	for node != nil {
		next := node.Next
		if node.Child != nil {
			child := node.Child
			childTail := flattenGetTail(child)

			node.Child = nil
			node.Next = child
			child.Prev = node

			childTail.Next = next
			if next != nil {
				next.Prev = childTail
			}

			tail = childTail
		} else {
			tail = node
		}

		node = next
	}

	return tail
}
