package l380

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
