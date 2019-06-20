package l3

func min(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

//Runtime: 4 ms, faster than 96.14% of Go online submissions for Longest Substring Without Repeating Characters.
//Memory Usage: 3.3 MB, less than 37.64% of Go online submissions for Longest Substring Without Repeating Characters.

// 设字符串 s 的长度为 N
// 时间复杂度 = O(2N) = O(N)
// 空间复杂度 = O(N) + O(122) = O(N) + C = O(N)
func lengthOfLongestSubstring(s string) int {
	var (
		length         int
		latestIndexMap = make([]int, 128) // ascii 码的范围是0~128
		key            = make([]int, len(s))
	)

	for j := 0; j < len(latestIndexMap); j++ {
		latestIndexMap[j] = -1
	}
	for i, char := range s {
		if i == 0 || latestIndexMap[char] == -1 {
			key[i] = key[max(0, i-1)] + 1
		} else {
			key[i] = min(key[i-1]+1, i-latestIndexMap[char])
		}
		latestIndexMap[char] = i

		length = max(length, key[i])
	}

	return length
}
