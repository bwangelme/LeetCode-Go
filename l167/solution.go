package l167

/*
N 表示 numbers 的长度
时间复杂度: O(N)
空间复杂度: O(1)
*/
func twoSum(numbers []int, target int) []int {
	var (
		i = 0
		j = len(numbers) - 1
	)

	for i < j && numbers[i]+numbers[j] != target {
		if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}

	return []int{i + 1, j + 1}
}
