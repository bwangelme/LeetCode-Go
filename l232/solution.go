package l232

/**
## 题目
请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：

实现 MyQueue 类：

void push(int x) 将元素 x 推到队列的末尾
int pop() 从队列的开头移除并返回元素
int peek() 返回队列开头的元素
boolean empty() 如果队列为空，返回 true ；否则，返回 false
说明：

你 只能 使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。

## 复杂度分析

### 时间复杂度

每次 Pop 和 Peek 的时候，会将 lstack 的数据移动到 rstack, 这个移动次数是可以均摊的，所以时间复杂度是 O(1)

### 空间复杂度

这个算法输入中没有 n, 整体申请的空间和队列的长度成线性关系，所以空间复杂度是O(n)
*/

type Stack[T any] []T

func NewStack[T any]() Stack[T] {
	return Stack[T]{}
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) Push(elem T) bool {
	*s = append(*s, elem)

	return true
}

func (s *Stack[T]) Pop() *T {
	stackLen := len(*s)
	if stackLen == 0 {
		return nil
	}

	// 注意: *s 必须扩起来，[] 的优先级高于 *
	top := (*s)[stackLen-1]
	*s = (*s)[:stackLen-1]

	return &top
}

func (s *Stack[T]) Peek() *T {
	stackLen := len(*s)
	if stackLen == 0 {
		return nil
	}
	top := (*s)[stackLen-1]
	return &top
}

type MyQueue struct {
	lStack Stack[int]
	rStack Stack[int]
}

func Constructor() MyQueue {
	return MyQueue{
		lStack: NewStack[int](),
		rStack: NewStack[int](),
	}
}

func (this *MyQueue) Push(x int) {
	this.lStack.Push(x)
}

func (this *MyQueue) move() {
	if this.rStack.Len() > 0 {
		return
	}

	for this.lStack.Len() > 0 {
		val := this.lStack.Pop()
		this.rStack.Push(*val)
	}
}

func (this *MyQueue) Pop() int {
	this.move()
	if this.rStack.Len() > 0 {
		return *this.rStack.Pop()
	}

	return -1
}

func (this *MyQueue) Peek() int {
	this.move()
	if this.rStack.Len() > 0 {
		return *this.rStack.Peek()
	}

	return -1
}

func (this *MyQueue) Empty() bool {
	return this.rStack.Len() == 0 && this.lStack.Len() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
