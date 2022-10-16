package main

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	slow = nums[slow]
	fast = nums[nums[fast]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	p, q := 0, slow
	for p != q {
		p = nums[p]
		q = nums[q]
	}
	return p
}
