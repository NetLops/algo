package main

type pair struct {
	val  int
	time int
}
type RecentCounter struct {
	times []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		times: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.times = append(this.times, t)
	//i :=
	left, right := 0, len(this.times)-1
	target := t - 3000
	for left < right {
		mid := left + (right-left)>>1
		if this.times[mid] < target {
			left = mid + 1
		} else if this.times[mid] == target {
			break
		} else {
			right = mid
		}
	}
	this.times = this.times[left-1:]
	return len(this.times)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
func main() {

}
