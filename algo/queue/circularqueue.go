package queue

//CircularQueue 循环队列
type CircularQueue[T any] struct {
	Data []T
	head int
	tail int
	size int
}

func NewCircularQueue[T any](capacity int) *CircularQueue[T] {
	// 有一个节点是哨兵节点，尾部要留一个节点的空间，用来判断列表是否满了
	return &CircularQueue[T]{
		Data: make([]T, capacity+1),
		size: capacity + 1,
		head: 0,
		tail: 0,
	}
}

func (q *CircularQueue[T]) Enqueue(data T) bool {
	if q.IsFull() {
		return false
	}

	q.Data[q.tail] = data
	q.tail = (q.tail + 1) % q.size
	return true
}

//Dequeue 从队头出队
// res 默认是 T 的0值，可能是 "", 0, nil 或空结构体
func (q *CircularQueue[T]) Dequeue() (res T) {
	if q.IsEmpty() {
		return
	}

	res = q.Data[q.head]
	q.head = (q.head + 1) % q.size
	return res
}

func (q *CircularQueue[T]) IsFull() bool {
	if (q.tail+1)%q.size == q.head {
		return true
	}

	return false
}

func (q *CircularQueue[T]) IsEmpty() bool {
	return q.tail == q.head
}

func (q *CircularQueue[T]) Size() int {
	return (q.tail + q.size - q.head) % q.size
}
