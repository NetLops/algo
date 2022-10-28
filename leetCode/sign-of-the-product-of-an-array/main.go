package main

func arraySign(nums []int) int {
	sign, i := 1, 0
	for ; i < len(nums); i++ {
		if nums[i] < 0 {
			sign = -sign
		} else if nums[i] > 0 {
			sign = sign
		} else {
			return 0
		}
	}

	return sign
}
