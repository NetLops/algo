package main

func exchange(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]%2 == 0 && nums[right]%2 == 1 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		} else if nums[left]%2 == 1 {
			left++
		} else if nums[right]%2 == 0 {
			right--
		}
	}
	return nums
}
func main() {

}
