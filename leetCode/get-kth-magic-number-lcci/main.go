package main

//
//func getKthMagicNumber(k int) int {
//	if k == 1 {
//		return 1
//	}
//	base := []int{3, 5, 7}
//	q := IntHeap{}
//	set := map[int]struct{}{} // 去重
//
//	heap.Push(&q, 1)
//	set[1] = struct{}{}
//	for len(q) != 0 {
//		temp := heap.Pop(&q).(int)
//		k--
//		if k == 0 {
//			return temp
//		}
//		for _, num := range base {
//			now := num * temp
//			if _, ok := set[now]; !ok {
//				heap.Push(&q, now)
//				set[now] = struct{}{}
//			}
//		}
//	}
//	return -1
//}
//
//type IntHeap []int
//
//func (h IntHeap) Len() int {
//	return len(h)
//}
//
//func (h IntHeap) Less(i, j int) bool {
//	if h[i] < h[j] {
//		return true
//	}
//	return false
//}
//
//func (h IntHeap) Swap(i, j int) {
//	h[i], h[j] = h[j], h[i]
//}
//
//func (h *IntHeap) Push(x interface{}) {
//	(*h) = append((*h), x.(int))
//}
//
//func (h *IntHeap) Pop() interface{} {
//	n := len((*h))
//	x := (*h)[n-1]
//	*h = (*h)[:n-1]
//	return x
//}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func getKthMagicNumber(k int) int {
	ans := make([]int, k+1)
	for i, i3, i5, i7 := 2, 1, 1, 1; i <= k; i++ {
		a, b, c := i3*3, i5*5, i7*7
		min := minInt(a, minInt(b, c))
		if min == a { // 三指针
			i3++
		}
		if min == b {
			i5++
		}
		if min == c {
			i7++
		}

		ans[i] = min
	}
	return ans[k]
}
func main() {

}
