package main

func productExceptSelf(nums []int) []int {
	res, p, q := make([]int, len(nums)), 1, 1
	for i := range res {
		res[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		res[i] *= p
		p *= nums[i]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= q
		q *= nums[i]
	}

	return res

}
