package l760

// 问题描述:
//
//      数组B是数组的异构数组，即A和B中的数字是相同的，但是它们的顺序是不一样的。
//
//      求一个数组P，令P[i] = j，其中A[i] == B[j]
//
//  如果A或者B中含有重复元素，那么返回的结果应该是多个的

// 求解思路
//
//  创建一个 Hash 数据结构bMap，key 是 B 中的每个数，Value 是这个数对应的索引列表，
// 因为B中可能存在着相同的数，所以 Value 应该是一个列表
//
//  遍历数组A，通过其值去查找bMap，将其索引添加到结果数组中
// TODO: 输出所有的结果

func anagramMappings(A []int, B []int) []int {
	if len(B) != len(A) {
		return make([]int, 0)
	}

	var bMap = make(map[int][]uint8)
	var answer = make([]int, len(B))

	for i := 0; i < len(B); i++ {
		valueList, exist := bMap[B[i]]
		if !exist {
			valueList = []uint8{}
		}

		bMap[B[i]] = append(valueList, uint8(i))
	}

	for i := 0; i < len(A); i++ {
		answer[i] = int(bMap[A[i]][0])
		bMap[A[i]] = bMap[A[i]][1:]
	}

	return answer
}
