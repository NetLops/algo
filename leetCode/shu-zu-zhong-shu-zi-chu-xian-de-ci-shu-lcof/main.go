package main

func singleNumbers(nums []int) []int {
	ret := 0
	// 找出 a ^ b
	for _, num := range nums {
		ret ^= num
	}

	d := 1 // 找出a 与 b 的 异
	for d&ret == 0 {
		d <<= 1
	}
	a, b := 0, 0
	// 其他都是两两抵消
	for _, num := range nums {
		if d&num != 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}
func main() {
	singleNumbers([]int{4, 1, 4, 6})

	//fmt.Println(17 ^ 3 ^ 1)
	//fmt.Println(1 ^ 3 ^ 17)
	//fmt.Println(1 ^ 17 ^ 3)
	//fmt.Println(1 & 1)
	//fmt.Println(1 & 3)
}
