package lcr41

type MovingAverage struct {
	size int
	nums []int
	sum  int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		size: size,
		nums: make([]int, 0),
		sum:  0,
	}
}

func (m *MovingAverage) Next(val int) float64 {
	m.nums = append(m.nums, val)
	m.sum += val
	if len(m.nums) > m.size {
		v := m.nums[0]
		m.nums = m.nums[1:]
		m.sum -= v
	}

	return float64(m.sum) / float64(len(m.nums))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
