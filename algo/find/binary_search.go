package find

// BinarySearchFirst
/**
返回 arr 中第一个 == num 的数字的索引

## 复杂度分析

最坏情况下，所有数字都相等，此时时间复杂度为 O(logN)

由于运算过程中没有额外申请和 len(arr) 有关的内存空间，所以空间复杂度是 O(1)
 **/
func BinarySearchFirst(arr []int64, num int64) int {
	var (
		start = 0
		end   = len(arr) - 1
	)

	for start <= end {
		mid := start + ((end - start) >> 1)
		midNum := arr[mid]
		if midNum > num {
			end = mid - 1
			continue
		} else if midNum < num {
			start = mid + 1
			continue
		}

		if num == midNum {
			// start 是从0一直往上加的，所以 start 最小是0
			// 因为 start 最小是0, mid = start + (end-start)/2, end-start 最小是0, 所以 mid 最小是0, 不会出现 -1
			if mid == 0 || arr[mid-1] != num {
				return mid
			} else {
				end = mid - 1
			}
		}
	}

	return -1
}

// BinarySearchFirstV2
/**
## 功能

返回 arr 中第一个 == num 的数字的索引

 **/
func BinarySearchFirstV2(arr []int64, num int64) int {
	var (
		low  = 0
		n    = len(arr)
		high = n - 1
	)

	/**
	## 对于 num 存在的情况
	1. for 循环结束后，low > high
	2. 由于判断条件 midNum >= num 先检查，所以循环结束后， arr[high] != num
	3. low > high, arr[high] != num, 所以 arr[low] 是第一个 num

	## 对于 arr 所有数比 num 小的情况
	1. for 循环结束后，low > high
	2. high 一次也没移动，low = n
	3. 此时 low 超出数组长度，arr[low] 会异常，所以要先判断 low < n

	## 对于 arr 所有数比 num 大的情况
	1. for 循环结束后，low > high
	2. low 一次也没移动，low = 0, high = -1
	3. 此时 arr[low] != num 不存在

	## 对于 num 处于 arr 中部但不存在的情况
	1. for 循环结束后，low > high
	2. low < n, arr[low] != num, return -1
	*/

	for low <= high {
		mid := low + ((high - low) >> 1)
		midNum := arr[mid]
		if midNum >= num {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if low < n && arr[low] == num {
		return low
	}

	return -1
}

// BinarySearchLast
/**
返回 arr 中最后一个 == num 的数字的索引

## 复杂度分析

最坏情况下，所有数字都相等，此时时间复杂度为 O(logN)

由于运算过程中没有额外申请和 len(arr) 有关的内存空间，所以空间复杂度是 O(1)
 **/
func BinarySearchLast(arr []int64, num int64) int {
	var (
		start = 0
		n     = len(arr)
		end   = n - 1
	)

	for start <= end {
		mid := start + ((end - start) >> 1)
		midNum := arr[mid]
		if midNum > num {
			end = mid - 1
			continue
		} else if midNum < num {
			start = mid + 1
			continue
		}

		if num == midNum {
			if mid == n-1 || arr[mid+1] != num {
				return mid
			} else {
				start = mid + 1
			}
		}
	}

	return -1
}

// BinarySearchLastV2
/**
## 功能

返回 arr 中最后一个 == num 的数字的索引

 **/
func BinarySearchLastV2(arr []int64, num int64) int {
	var (
		low  = 0
		n    = len(arr)
		high = n - 1
	)

	/**
	1. for 循环结束后, low > high
	2. 由于 midNum <= num 条件先执行，for 结束后, low 指向第一个 > n 的数，high 指向 low 的前一位
	3. 如果 high 的索引在数组中，且 arr[high] == num, 那么 high 就是 num 在数组中最后一位
	*/
	for low <= high {
		mid := low + ((high - low) >> 1)
		midNum := arr[mid]
		if midNum <= num {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if high > 0 && arr[high] == num {
		return high
	}

	return -1
}
