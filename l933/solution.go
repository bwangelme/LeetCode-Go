package l933

/*
## 题目

写一个 RecentCounter 类来计算特定时间范围内最近的请求。

请你实现 RecentCounter 类：

RecentCounter() 初始化计数器，请求数为 0 。
int ping(int t) 在时间 t 添加一个新请求，其中 t 表示以毫秒为单位的某个时间，并返回过去 3000 毫秒内发生的所有请求数（包括新请求）。确切地说，返回在 [t-3000, t] 内发生的请求数。
保证 每次对 ping 的调用都使用比之前更大的 t 值。

提示：

1 <= t <= 109
保证每次对 ping 调用所使用的 t 值都 严格递增
至多调用 ping 方法 104 次

## 时间复杂度

假设窗口大小为 w 毫秒

空间复杂度 O(w)
	times 数组大小和 w 有关
时间复杂度 O(w)
	for 循环移除次数上线是 w

但因为 w 是常数 3000, 因此时间空间复杂度都是 O(1)

*/

type RecentCounter struct {
	windowSize int
	times      []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		windowSize: 3000,
		times:      make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.times = append(this.times, t)
	for {
		v := this.times[0]
		if v+this.windowSize < t {
			this.times = this.times[1:]
		} else {
			break
		}
	}

	return len(this.times)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
