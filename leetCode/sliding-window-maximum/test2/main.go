package main

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	q := []int{}
	push := func(i int) {
		// 问题点：q时间问题 比如 (i - k,  i] 之间 存在 中间的数大于后面部分数
		// 能保证]
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}
	ans := make([]int, n-k+1)
	t := 1
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i-k {
			q = q[1:]
		}
		ans[t] = nums[q[0]]
		t++
	}
	return ans
}
func main() {

}
