package main

//
//func findUnsortedSubarray(nums []int) int {
//	arr := make([]int, len(nums))
//	copy(arr, nums)
//	sort.Ints(arr)
//	i, j := 0, len(nums)-1
//	for i <= j && nums[i] == arr[i] {
//		i++
//	}
//	for i <= j && nums[j] == arr[j] {
//		j--
//	}
//	return j - i + 1
//}
func findUnsortedSubarray(nums []int) int {
	MIN := -100010
	MAX := 100010
	n := len(nums)
	i, j := 0, n-1
	for i < j && nums[i] <= nums[i+1] {
		i++
	}
	for i < j && nums[j] >= nums[j-1] {
		j--
	}
	l, r := i, j
	mins, maxs := nums[i], nums[j]
	for a := l; a <= r; a++ {
		if nums[a] < mins {
			for i >= 0 && nums[i] > nums[a] {
				i--
			}
			if i >= 0 {
				mins = nums[i]
			} else {
				mins = MIN
			}
		}
		if nums[a] > maxs {
			for j < n && nums[j] < nums[a] {
				j++
			}
			if j < n {
				maxs = nums[j]
			} else {
				maxs = MAX
			}
		}
	}
	if j == i {
		return 0
	} else {
		return j - 1 - (i + 1) + 1
	}
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	ans := 0
	for i < j {
		temp := (j - i) * Min(height[i], height[j])
		if ans < temp {
			ans = temp
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return ans
}
