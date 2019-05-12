package l1

func twoSum(nums []int, target int) []int {
	buf_map := make(map[int]int)

	for idx, num := range nums {
		if _, ok := buf_map[num]; ok {
			return []int{
				buf_map[num], idx,
			}
		} else {
			buf_map[target-num] = idx
		}
	}

	panic("Invalid Input")
}
