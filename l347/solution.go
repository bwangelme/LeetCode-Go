package l347

import "container/heap"

type num2Count struct {
	Num   int
	Count int
}

type num2CountSlice []num2Count

func (n *num2CountSlice) Len() int           { return len(*n) }
func (n *num2CountSlice) Less(i, j int) bool { return (*n)[i].Count < (*n)[j].Count }
func (n *num2CountSlice) Swap(i, j int)      { (*n)[i], (*n)[j] = (*n)[j], (*n)[i] }
func (n *num2CountSlice) Peek() any          { return (*n)[0] }
func (n *num2CountSlice) Push(x any) {
	*n = append(*n, x.(num2Count))
}

func (n *num2CountSlice) Pop() (v any) {
	*n, v = (*n)[:n.Len()-1], (*n)[n.Len()-1]
	return v
}

func topKFrequent(nums []int, k int) []int {
	var numToCount = make(map[int]int, 0)
	for _, num := range nums {
		v, _ := numToCount[num]
		numToCount[num] = v + 1
	}
	var minHeap = &num2CountSlice{}

	heap.Init(minHeap)
	for num, count := range numToCount {
		item := num2Count{
			Num:   num,
			Count: count,
		}
		if minHeap.Len() < k {
			heap.Push(minHeap, item)
		} else {
			top := minHeap.Peek().(num2Count)
			if count > top.Count {
				heap.Pop(minHeap)
				heap.Push(minHeap, item)
			}
		}
	}

	var result = make([]int, 0)
	for _, v := range *minHeap {
		result = append(result, v.Num)
	}

	return result
}
