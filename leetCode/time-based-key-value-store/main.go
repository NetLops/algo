package main

type node struct {
	timeStamp int
	value     string
}

type TimeMap struct {
	m map[string][]node
}

func Constructor() TimeMap {
	return TimeMap{
		m: map[string][]node{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.m[key]; !ok {
		this.m[key] = []node{}
	}

	this.m[key] = append(this.m[key], node{
		timeStamp: timestamp,
		value:     value,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	// 存在否
	if _, ok := this.m[key]; !ok {
		return ""
	}

	if len(this.m[key]) == 0 {
		return ""
	}
	mnode := this.m[key]
	l, r := 0, len(mnode) // r == 数组长度 让 l 偏右一些
	for l < r {
		mid := l + (r-l)>>1
		if timestamp >= mnode[mid].timeStamp {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if l > 0 {
		return this.m[key][l-1].value
	}
	return ""
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
func main() {

}
