package l287

/*
## 题目

给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。

假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。

你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。

输入：nums = [1,3,4,2,2]
输出：2

输入：nums = [3,1,3,4,2]
输出：3

提示：

1 <= n <= 10^5
nums.length == n + 1
1 <= nums[i] <= n
nums 中 只有一个整数 出现 两次或多次 ，其余整数均只出现 一次

进阶：

如何证明 nums 中至少存在一个重复的数字?
你可以设计一个线性级时间复杂度 O(n) 的解决方案吗？

## 思路

因为

- 1 <= nums[i] <= n
nums.length == n + 1

我们可以将 nums[i] 作为下一个节点的索引，这样可以将数组转换成一个链表，而判断数组中是否有重复元素的办法，就是判断对应的链表中是否有环

例如 1,3,4,2,2 可以转换成

0 -> 1 -> 3 -> 2 -> 4
               |    |
               | <- |

查找数组中重复元素，就是查找链表中环的入口

链表的迭代办法

p = nums[p]

判断链表有环，我们可以用快慢双指针法

## 复杂度分析

- 空间复杂度

我们只申请了四个指向链表的指针，所以时间复杂度是 O(1)

- 时间复杂度

我们在链表上迭代的次数小于 2N, 所以时间复杂度是 O(N)
*/

func findDuplicate(nums []int) int {
	var (
		slow = nums[0]
		fast = nums[nums[0]]
	)

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	var (
		pre1 = 0
		pre2 = slow
	)

	for pre1 != pre2 {
		pre1 = nums[pre1]
		pre2 = nums[pre2]
	}

	return pre1
}
