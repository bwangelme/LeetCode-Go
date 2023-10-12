package l81

/*
## 题目

已知存在一个按非降序排列的整数数组 nums ，数组中的值不必互不相同。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转 ，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
例如， [0,1,2,4,4,4,5,6,6,7] 在下标 5 处经旋转后可能变为 [4,5,6,6,7,0,1,2,4,4] 。

给你 旋转后 的数组 nums 和一个整数 target ，请你编写一个函数来判断给定的目标值是否存在于数组中。如果 nums 中存在这个目标值 target ，则返回 true ，否则返回 false 。

你必须尽可能减少整个操作步骤。

## 分析

nums[mid]与nums[l]比较，有三种情况：

1.nums[mid] > nums[l],则表明[l,mid]是有序的,判断target是否位于有序区间[l,mid-1]中；否则target位于[mid,r]中

2.nums[mid] < nums[l]，则表明[mid,r]是有序的,判断target是否位于有序区间[mid+1,r]中；否则target位于[l,mid]中

3.nums[mid] == nums[l]
- 如果nums[l]!=target，我们可以排除左边界l (排除了一个数字);
- 如果nums[l]==target,那么nums[mid]==target，也可以排除左边界l, (前面已经判断过 nums[mid] == target, 所以这个情况不可能执行到)
- 所以这种情况下让l++即可

## 复杂度分析

- 时间复杂度

最坏情况是 nums 中所有数字都相等，且不等于 target, 此时复杂度为 O(n)

- 空间复杂度

没有根据 nums 的长度 N 申请新的空间，所以空间复杂度是 O(1)
*/
func search(nums []int, target int) bool {
	var (
		left  = 0
		right = len(nums) - 1
	)

	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] == target {
			return true
		}

		if nums[mid] > nums[left] {
			// 1.
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				// left --> mid 是有序的，但 left --> mid+1 不一定有序，所以这里 left 只能移动到 mid
				left = mid
			}
		} else if nums[mid] < nums[left] {
			// 2.
			if target <= nums[right] && target > nums[mid] {
				left = mid + 1
			} else {
				// 同理，只能保证 mid --> right 有序，不能保证 mid-1 到 right 有序
				right = mid
			}
		} else {
			// 3.
			left++
		}
	}
	return false
}
