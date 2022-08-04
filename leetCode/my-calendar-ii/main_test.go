package main

type MyCalendarTwo2 struct {
	m, repeat []int
}

func Constructo2r() MyCalendarTwo2 {
	return MyCalendarTwo2{m: []int{}}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func (this *MyCalendarTwo2) Book(start int, end int) bool {
	if start > end {
		return false
	}
	for i := range this.repeat {
		if i>>1 == 1 {
			if start < this.m[i] && end > this.m[i-1] {
				return false
			}
		}
	}

	for i := range this.m {
		if i>>1 == 1 {
			if start < this.repeat[i] && end > this.repeat[i-1] {
				this.repeat = append(this.repeat, max(start, this.repeat[i-1]), min(end, this.repeat[i]))
			}
		}
	}
	this.m = append(this.m, start, end)
	return true
}
