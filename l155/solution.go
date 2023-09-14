package l155

import "math"

/**
## 题目

设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

实现 MinStack 类:

MinStack() 初始化堆栈对象。
void push(int val) 将元素val推入堆栈。
void pop() 删除堆栈顶部的元素。
int top() 获取堆栈顶部的元素。
int getMin() 获取堆栈中的最小元素。

## 解题思路

创建栈的时候用了分治的思想

1. head.min 保存的是整个栈中最小的元素, 下面的每个节点, 表示如果以他为 head, head.min 是整个栈的最小值
2. len(minStack) == 1 的时候，head.min == head.val

Push 的时候，val < head.min, 则 newHead.min = val ，否则 newHead.min = head.min

## 复杂度分析

设栈的大小为 N

Push, Pop, Top, GetMin, 每个操作的时间复杂度都是 O(1)

空间复杂度是 O(N)

*/

type Node struct {
	val  int
	min  int
	next *Node
}

type MinStack struct {
	head *Node
}

func Constructor() MinStack {
	return MinStack{
		head: nil,
	}
}

func (this *MinStack) Push(val int) {
	if this.head == nil {
		this.head = &Node{
			val: val,
			min: val,
		}
	} else {
		this.head = &Node{
			val:  val,
			min:  int(math.Min(float64(val), float64(this.head.min))),
			next: this.head,
		}
	}
}

func (this *MinStack) Pop() {
	if this.head == nil {
		return
	}
	this.head = this.head.next
}

func (this *MinStack) Top() int {
	if this.head == nil {
		return 0
	}

	return this.head.val
}

func (this *MinStack) GetMin() int {
	if this.head == nil {
		return 0
	}
	return this.head.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
