package l146

/*
## 题目

请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。



示例：

输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4


提示：

1 <= capacity <= 3000
0 <= key <= 10000
0 <= value <= 105
最多调用 2 * 10^5 次 get 和 put

## 复杂度分析

n 表示链表的长度

Get: 时间复杂度 O(1)
Put: 时间复杂度 O(1)

空间复杂度: O(n)
*/

type Node struct {
	prev  *Node
	next  *Node
	key   int
	value int
}

type LRUCache struct {
	head     *Node
	tail     *Node
	table    map[int]*Node
	capacity int
}

func Constructor(capacity int) LRUCache {
	head := &Node{key: -1, value: -1}
	tail := &Node{key: -1, value: -1}

	head.prev = nil
	head.next = tail
	tail.prev = head
	tail.next = nil

	l := &LRUCache{
		head:     head,
		tail:     tail,
		table:    make(map[int]*Node),
		capacity: capacity,
	}

	return *l
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.table[key]
	if !ok {
		return -1
	}

	this.moveToTail(node, node.value)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.table[key]
	if ok {
		this.moveToTail(node, value)
	} else {
		if len(this.table) == this.capacity {
			toBeDeleted := this.head.next
			deleteNode(toBeDeleted)
			delete(this.table, toBeDeleted.key)
		}

		node := &Node{key: key, value: value}
		this.insertToTail(node)
		this.table[node.key] = node
	}
}

func (this *LRUCache) insertToTail(node *Node) {
	this.tail.prev.next = node
	node.prev = this.tail.prev
	node.next = this.tail
	this.tail.prev = node
}

func (this *LRUCache) moveToTail(node *Node, newValue int) {
	deleteNode(node)
	node.value = newValue
	this.insertToTail(node)
}

func deleteNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
