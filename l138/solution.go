package l138

/*
## 题目

给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，该指针可以指向链表中的任何节点或空节点。

构造这个链表的 深拷贝。 深拷贝应该正好由 n 个 全新 节点组成，其中每个新节点的值都设为其对应的原节点的值。新节点的 next 指针和 random 指针也都应指向复制链表中的新节点，并使原链表和复制链表中的这些指针能够表示相同的链表状态。复制链表中的指针都不应指向原链表中的节点 。

例如，如果原链表中有 X 和 Y 两个节点，其中 X.random --> Y 。那么在复制链表中对应的两个节点 x 和 y ，同样有 x.random --> y 。

返回复制链表的头节点。

用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：

    val：一个表示 Node.val 的整数。
    random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  null 。

你的代码 只 接受原链表的头节点 head 作为传入参数。

提示：

    0 <= n <= 1000
    -10^4 <= Node.val <= 10^4
    Node.random 为 null 或指向链表中的节点。

## 解题思路

假设存在链表 A -> B -> C

1. 将链表每个节点都复制一份，变成

A -> A' -> B -> B' -> C -> C'

复制节点中 random 和 原节点的 random 指向相同

2. 将 random 指针修正

遍历所有节点，

复制节点.random = 原来节点.random.next

3. 还原链表

将 A -> A' -> B -> B' -> C -> C'

还原成 A -> B -> C 和 A' -> B' -> C'

## 复杂度分析

时间复杂度, 循环了三次，遍历了 5N 个节点，时间复杂度是 O(N)
空间复杂度是 O(1)
*/

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func printList(head *Node) {
	fmt.Println("--------------")
	for head != nil {
		fmt.Printf("val: %v", head.Val)
		if head.Random != nil {
			fmt.Printf(" random: %v\n", head.Random.Val)
		} else {
			fmt.Printf(" random: nil\n")
		}
		head = head.Next
	}
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	var current = head
	for current != nil {
		copyNode := &Node{
			Val:    current.Val,
			Next:   current.Next,
			Random: current.Random,
		}
		current.Next = copyNode
		current = copyNode.Next
	}

	current = head
	for current != nil {
		copyNode := current.Next
		if current.Random != nil {
			copyNode.Random = current.Random.Next
		}
		current = copyNode.Next
	}

	current = head
	res := head.Next
	for current != nil {
		copyNode := current.Next
		current.Next = copyNode.Next

		if copyNode.Next != nil {
			copyNode.Next = copyNode.Next.Next
		}

		current = current.Next
	}

	return res
}
