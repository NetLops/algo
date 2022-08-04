package main

//
//func maxSubArray(nums []int) int {
//	p := make([]int, len(nums))
//	// 查找每行的最大值
//	p[0] = nums[0]
//	max := p[0]
//	for i := 1; i < len(nums); i++ {
//		if p[i-1] < nums[i] && p[i-1] >= 0 {
//			p[i] = nums[i]
//		} else {
//			p[i] = p[i-1] + nums[i]
//		}
//		if max < p[i] {
//			max = p[i]
//		}
//	}
//	return max
//}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func maxSubArray(nums []int) int {
	p := make([]int, len(nums))
	// 查找每行的最大值
	p[0] = nums[0]
	max := p[0]
	for i := 1; i < len(nums); i++ {
		p[i] = maxInt(p[i-1]+nums[i], nums[i])
		max = maxInt(max, p[i])
	}
	return max
}

func main() {

}
