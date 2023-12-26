package l338

func CountBits(num int) []int {
	var result []int
	var bitCount int

	result = append(result, 0)
	for i := 1; i <= num; i++ {
		bitCount = result[i>>1] + i&1
		result = append(result, bitCount)
	}

	return result
}

func CountBitsV2(num int) []int {
	var result = make([]int, num+1)
	for i := 1; i <= num; i++ {
		result[i] = result[i&(i-1)] + 1
	}

	return result
}
