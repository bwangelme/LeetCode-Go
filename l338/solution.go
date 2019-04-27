package l338

func CountBits(num int) []int {
	var result []int
	var bitCount int

	result = append(result, 0)
	for i := 1; i <= num; i++ {
		if i&1 == 1 {
			bitCount = result[i>>1] + 1
		} else {
			bitCount = result[i>>1]
		}

		result = append(result, bitCount)
	}

	return result
}
