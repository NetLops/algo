package main

import "sort"

func advantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	ans := make([]int, n)
	sort.Ints(nums1)
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = i
	}
	sort.Slice(ids, func(i, j int) bool {
		return nums2[ids[i]] < nums2[ids[j]]
	})
	left, right := 0, n-1
	for _, num := range nums1 {
		if num > nums2[ids[left]] {
			ans[ids[left]] = num // 用下等马比下等马
			left++
		} else { // 下等马比上等马
			ans[ids[right]] = num
			right--
		}
	}
	return ans
}
