package main

func sortColors(nums []int) {
	i := 0
	for k := 0; k < 3; k++ {
		for j := 0; j < len(nums); j++ {
			if nums[j] == k {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
		}
	}

}
