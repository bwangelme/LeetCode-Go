package l215

func findKthLargest(nums []int, k int) int {
	var (
		start  = 0
		end    = len(nums) - 1
		target = k - 1
	)
	for {
		idx := partitionV2(nums, start, end)
		if idx == target {
			return nums[idx]
		}

		if idx < target {
			start = idx + 1
		} else {
			end = idx - 1
		}
	}
}

func partitionV2(arr []int, p, r int) int {
	pivot := arr[r]
	i := p
	for j := p; j <= r-1; j++ {
		if arr[j] > pivot {
			if i != j {
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}

			i++
		}
	}

	tmp := arr[i]
	arr[i] = arr[r]
	arr[r] = tmp

	return i
}
