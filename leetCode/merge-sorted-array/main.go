package main

//func merge(nums1 []int, m int, nums2 []int, n int) {
//	for i := 0; i < n; i++ {
//		nums1[i+m] = nums2[i]
//	}
//	sort.Ints(nums1)
//}
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, tail := m-1, m-1, m+n-1
	for ; p1 > 0 && p2 > 0 && tail > 0; tail-- {
		cur := 0
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] < nums2[p2] {
			cur = nums2[p2]
			p2--
		} else {
			cur = nums1[p1]
			p1--
			p2--
		}
		nums1[tail] = cur
	}
}
func main() {

}
