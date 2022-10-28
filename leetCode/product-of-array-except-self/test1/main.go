package main

func productExceptSelf(nums []int) []int {
	p, q := 1, 1
	i, n := 0, len(nums)
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		ans[i] = 1
	}
	for i = 0; i < n; i++ {
		ans[i] *= p
		p *= nums[i]
	}

	for i = n - 1; i >= 0; i-- {
		ans[i] *= q
		q *= nums[i]
	}

	return ans
}
