package l729

import "github.com/emirpasic/gods/maps/treemap"

type MyCalendar struct {
	tree *treemap.Map
}

func Constructor() MyCalendar {
	return MyCalendar{
		tree: treemap.NewWithIntComparator(),
	}
}

func (this *MyCalendar) Book(startTime int, endTime int) bool {
	_, entryEndI := this.tree.Floor(startTime)
	if entryEndI != nil {
		entryEnd := entryEndI.(int)
		if entryEnd > startTime {
			return false
		}
	}

	entryStartI, _ := this.tree.Ceiling(startTime)
	if entryStartI != nil {
		entryStart := entryStartI.(int)
		if entryStart < endTime {
			return false
		}
	}

	this.tree.Put(startTime, endTime)

	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */
