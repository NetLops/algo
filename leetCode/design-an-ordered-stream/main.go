package main

type OrderedStream struct {
	m   map[int]string
	n   int
	ptr int
}

func (o *OrderedStream) display(id int) (s []string) {
	if id != o.ptr {
		return
	}
	for i := o.ptr; i <= o.n; i++ {
		if o.m[i] != "" {
			s = append(s, o.m[i])
		} else {
			o.ptr = i
			return
		}
	}
	return
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		m:   map[int]string{},
		n:   n,
		ptr: 1,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.m[idKey] = value
	return this.display(idKey)
}
func main() {

}
