package main

type Node struct {
	k, v string
	t    int
}
type TimeMap struct {
	m map[string][]Node
}

func Constructor() TimeMap {
	return TimeMap{
		m: map[string][]Node{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.m[key]; !ok {
		this.m[key] = []Node{}
	}
	this.m[key] = append(this.m[key], Node{key, value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, ok := this.m[key]; !ok {
		this.m[key] = []Node{}
	}
	if len(this.m[key]) == 0 {
		return ""
	}
	l, r := 0, len(this.m[key])-1
	for l < r {
		mid := (l + r + 1) >> 1
		if this.m[key][mid].t <= timestamp {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if this.m[key][r].t <= timestamp {
		return this.m[key][r].v
	} else {
		return ""
	}
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
func main() {

}
