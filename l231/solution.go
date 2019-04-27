package l231

func isPowerOfTwo(n int) bool {
	return n & (n-1) == 0 && n >= 1
}
