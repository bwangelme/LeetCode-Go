package l318

func maxProduct(words []string) int {
	var flags = make([]uint32, len(words))

	for i, word := range words {
		for _, c := range word {
			flags[i] |= 1 << (c - 'a')
		}
	}

	var result = 0
	for i, x := range flags {
		for j, y := range flags[:i] {
			if x&y == 0 && len(words[i])*len(words[j]) > result {
				result = len(words[i]) * len(words[j])
			}
		}
	}

	return result
}
