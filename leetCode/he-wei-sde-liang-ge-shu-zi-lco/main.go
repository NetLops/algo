package main

//func twoSum(nums []int, target int) []int {
//
//	p := make(map[int]int)
//	for i := 0; nums[i] < target; i++ {
//		if k, ok := p[target-nums[i]]; ok {
//			return []int{nums[k], nums[i]}
//		} else {
//			p[nums[i]] = i
//		}
//	}
//	return []int{}
//}

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			return []int{nums[left], nums[right]}
		}
	}
	return []int{}
}
func main() {
	twoSum([]int{10, 26, 30, 31, 47, 60}, 40)
}
