package main

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func minSwap(nums1 []int, nums2 []int) int {
	n := len(nums1)
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, 2)
	}
	for i := 1; i < n; i++ {
		f[i][0], f[i][1] = n+10, n+10
	}
	f[0][1] = 1
	for i := 1; i < n; i++ {
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			f[i][0] = f[i-1][0]
			f[i][1] = f[i-1][1] + 1
		}
		if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
			f[i][0] = min(f[i][0], f[i-1][1])
			f[i][1] = min(f[i][1], f[i-1][0]+1)
		}
	}
	return min(f[n-1][0], f[n-1][1])
}
