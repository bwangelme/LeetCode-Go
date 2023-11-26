package eight_queen

import "fmt"

// Per 一种排列
type Per []int

func Permutation(N int) []Per {
	var target = make(Per, N)
	var result = make([]Per, 0)
	var choices []int
	for i := 1; i <= N; i++ {
		choices = append(choices, i)
	}
	permutation(choices, 0, &target, &result)

	return result
}

func permutation(choices []int, index int, target *Per, result *[]Per) {
	if len(choices) == 1 {
		(*target)[index] = choices[0]
		storeTarget := make(Per, len(*target))
		copy(storeTarget, *target)
		*result = append(*result, storeTarget)
		return
	}

	for idx, num := range choices {
		cpChoices := make([]int, len(choices))
		copy(cpChoices, choices)
		newChoices := append(cpChoices[0:idx], cpChoices[idx+1:]...)
		(*target)[index] = num
		permutation(newChoices, index+1, target, result)
	}
}

func tmp() {
	choices := []int{1, 2, 3}
	a := choices[0:1]
	b := choices[2:]
	// 注意 a b choices 都是基于同一个数组，将 a b 合并到一起会修改底层的数组
	newChoices := append(a, b...)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(newChoices)
	fmt.Println(choices)

	// choices == [1, 3, 3]
}
