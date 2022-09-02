package main

import "fmt"

func main() {
	var n int
	ans := 0
	_, _ = fmt.Scanln(&n)
	nums := make([]int, n)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &nums[i])
	}
	//fmt.Scanln()
	//fmt.Println(nums, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]-nums[j] == 2*nums[j]-nums[k] {
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
}
