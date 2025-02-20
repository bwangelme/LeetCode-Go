package l88

/*
merge

## 复杂度分析

时间复杂度: 便利了 nums1 和 nums2 两个数组，最后再循环复制了一遍

因此时间复杂度是 O(m+n)

空间复杂度

使用了一个额外的数组，res，因此空间复杂度是 O(m+n)
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	res := make([]int, m+n)
	var (
		mIdx   = 0
		nIdx   = 0
		resIdx = 0
	)

	for mIdx < m && nIdx < n {
		mVal := nums1[mIdx]
		nVal := nums2[nIdx]
		if mVal <= nVal {
			res[resIdx] = mVal
			mIdx++
		} else {
			res[resIdx] = nVal
			nIdx++
		}

		resIdx++
	}

	if mIdx >= m {
		for nIdx < n {
			res[resIdx] = nums2[nIdx]
			resIdx++
			nIdx++
		}
	} else {
		for mIdx < m {
			res[resIdx] = nums1[mIdx]
			resIdx++
			mIdx++
		}
	}

	copy(nums1[:], res[:])
}
