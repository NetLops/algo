package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums2); i++ {
		m[nums2[i]] = i
	}

	ans := make([]int, len(nums1))
	for i := range ans {
		ans[i] = -1
	}
	for i := range nums1 {
		for j := m[nums1[i]]; j < len(nums2); j++ {
			if nums2[j] > nums1[i] {
				ans[i] = nums2[j]
				break
			}
		}
	}
	return ans
}
func main() {

}
