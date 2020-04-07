package l341

/*
	这道题在 LeetCode 上编写，无法在本地编写
 */

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {}


type NestedIterator struct {
	numbers []int
}

func parseNestedIntegers(nestedList []*NestedInteger, it *NestedIterator) {
	for _, nestedInt := range nestedList {
		if nestedInt.IsInteger() {
			it.AddNumber(nestedInt.GetInteger())
		} else {
			list := nestedInt.GetList()
			parseNestedIntegers(list, it)
		}
	}
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	var res = NestedIterator{
		numbers: make([]int, 0),
	}

	var it = &res
	parseNestedIntegers(nestedList, it)
	return it
}

func (this *NestedIterator) AddNumber(number int) {
	this.numbers = append(this.numbers, number)
}

func (this *NestedIterator) Next() int {
	number := this.numbers[0]
	this.numbers = this.numbers[1:]
	return number
}

func (this *NestedIterator) HasNext() bool {
	return len(this.numbers) != 0
}

