package l55

/*
## 题目

给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。

1 <= nums.length <= 10^4
0 <= nums[i] <= 10^5

*/

/*
canJump

## 解题思路

用 i 从后向前遍历数组，判断 nums[i] 是否是0,

	如果不是0
		一直往前遍历, 直到 i < 0
	如果是 0
		用 j 遍历从 i-1 到 0 的所有数字
		判断 i 当前是否是最后一位
			如果是最后一位
				寻找 i 前面是否有位置可以超过 i (nums[j]+j > i)
			如果不是最后一位
				寻找 i 前面是否有位置可以到达 i (nums[j]+j >= i)
		用 j 算出 i 的新位置，如果新位置的 nums[i] 依然是 0
			因为遇到了两个 0，return false

如果 i 遍历到了 0, 那么说明从后往前可以到达，return true

## 复杂度分析

n 表示 len(nums)

### 时间复杂度

无论是 2, 3, 1, 1, 4 这种情况
还是 0, 1, 0, 1, 0 这种情况，实际循环的次数都是 N

因此时间复杂度是 O(N)

### 空间复杂度

没有申请与 N 相关的数组，因此空间复杂度是 O(1)
*/
func canJump(nums []int) bool {
	var n = len(nums)
	if n <= 1 {
		return true
	}

	for i := n - 1; i >= 0; {
		if nums[i] != 0 {
			i--
			continue
		} else {
			for j := i - 1; j >= 0; {
				if nums[j]+j > i {
					i = j
					break
				}

				if i == n-1 && nums[j]+j == i {
					i = j
					break
				}
				j--
			}

			if nums[i] == 0 {
				return false
			}
		}
	}

	return true
}
