package l347

import (
	"container/heap"

	"github.com/bwangelme/LeetCode-Go/lt"
)

type num2Count struct {
	Num   int
	Count int
}

func topKFrequent(nums []int, k int) []int {
	var numToCount = make(map[int]int, 0)
	for _, num := range nums {
		v, _ := numToCount[num]
		numToCount[num] = v + 1
	}
	var minHeap = lt.NewHeap[*num2Count](func(v1 *num2Count, v2 *num2Count) bool {
		return v1.Count < v2.Count
	})

	heap.Init(minHeap)
	for num, count := range numToCount {
		item := num2Count{
			Num:   num,
			Count: count,
		}
		if minHeap.Len() < k {
			minHeap.Offer(&item)
		} else {
			top := minHeap.Peek()
			if count > top.Count {
				minHeap.Poll()
				minHeap.Offer(&item)
			}
		}
	}

	var result = make([]int, 0)
	for _, v := range minHeap.ToArray() {
		result = append(result, v.Num)
	}

	return result
}
