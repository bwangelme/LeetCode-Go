package l380

/*
## 题目

实现RandomizedSet 类：

RandomizedSet() 初始化 RandomizedSet 对象
bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。



示例：

输入
["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
输出
[null, true, false, true, 2, true, false, 2]

解释
RandomizedSet randomizedSet = new RandomizedSet();
randomizedSet.insert(1); // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
randomizedSet.remove(2); // 返回 false ，表示集合中不存在 2 。
randomizedSet.insert(2); // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
randomizedSet.getRandom(); // getRandom 应随机返回 1 或 2 。
randomizedSet.remove(1); // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
randomizedSet.insert(2); // 2 已在集合中，所以返回 false 。
randomizedSet.getRandom(); // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。


提示：

-2^31 <= val <= 2^31 - 1
最多调用 insert、remove 和 getRandom 函数 2 * 10^5 次
在调用 getRandom 方法时，数据结构中 至少存在一个 元素。

## 复杂度分析

三个函数的时间复杂度都是 O(1)
空间复杂度是 O(n)
*/

import "math/rand"

type RandomizedSet struct {
	nums           []int
	numsToLocation map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums:           make([]int, 0),
		numsToLocation: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.numsToLocation[val]
	if ok {
		return false
	}

	location := len(this.nums)
	this.nums = append(this.nums, val)
	this.numsToLocation[val] = location

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.numsToLocation[val]
	if !ok {
		return false
	}

	last := len(this.nums) - 1
	lastVal := this.nums[last]
	this.nums[idx] = lastVal
	this.numsToLocation[lastVal] = idx
	this.nums = this.nums[0:last]
	delete(this.numsToLocation, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
