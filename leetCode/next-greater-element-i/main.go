package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	//pre := nums2[0]
	m := map[int]int{}
	stack := []int{} // 单调栈
	for i := len(nums2) - 1; i >= 0; i-- {

		for len(stack) > 0 && stack[len(stack)-1] <= nums2[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			m[stack[len(stack)-1]] = nums2[i]
		} else {
			m[nums2[i]] = -1
		}
		stack = append(stack, nums2[i])
	}
	ans := make([]int, 0, len(nums1))
	for _, v := range nums1 {
		ans = append(ans, m[v])
	}
	return ans
}
func main() {

}
