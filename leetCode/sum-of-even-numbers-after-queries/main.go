package main

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	ans := []int{}
	sum := 0
	for i := range nums {
		if nums[i]%2 == 0 {
			sum += nums[i]
		}
	}

	for _, query := range queries {
		// 减少 默认位置为偶数的影响
		if nums[query[1]]%2 == 0 {
			sum -= nums[query[1]]
		}
		nums[query[1]] += query[0]
		// 在判断是奇偶
		if nums[query[1]]%2 == 0 {
			sum += nums[query[1]]
		}
		ans = append(ans, sum)
	}
	return ans
}
func main() {

}
