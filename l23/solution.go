package l23

type ListNode struct {
	Val  int
	Next *ListNode
}

type OrderMap struct {
	order []int
	m     map[int][]*ListNode
}

func (o *OrderMap) Set(node *ListNode) {
	v, ok := o.m[node.Val]
	if !ok {
		val := node.Val + 10000
		o.order[val] = 1
	}
	v = append(v, node)
	o.m[node.Val] = v
}

func (o *OrderMap) MergeLoop() []*ListNode {
	var res = make([]*ListNode, 0)
	for idx, exist := range o.order {
		if exist == 1 {
			val := idx - 10000
			ls, _ := o.m[val]
			for _, node := range ls {
				res = append(res, node)
			}
		}
	}

	return res
}

func NewOrderMap() *OrderMap {
	return &OrderMap{
		order: make([]int, 20000+2),
		m:     make(map[int][]*ListNode),
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	var resMap = NewOrderMap()
	for _, list := range lists {
		current := list
		for current != nil {
			resMap.Set(current)
			current = current.Next
		}
	}

	var (
		res     = &ListNode{}
		current = res
	)
	for _, node := range resMap.MergeLoop() {
		current.Next = node
		current = current.Next
	}

	return res.Next
}
