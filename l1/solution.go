package l1

func twoSum(nums []int, target int) []int {
	bufMap := make(map[int]int)

	for idx, num := range nums {
		if _, ok := bufMap[num]; ok {
			return []int{
				bufMap[num], idx,
			}
		} else {
			bufMap[target-num] = idx
		}
	}

	panic("Invalid Input")
}
