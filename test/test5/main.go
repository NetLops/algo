package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scanln(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &nums[i])
	}
	//fmt.Println()
	for _, v := range getAns(n, nums) {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func getAns(length int, nums []int) (ans []int) {
	for i := length - 1; i >= 0; i-- {
		ans = getPos(nums[i], ans)
	}
	return
}

func getPos(num int, nums []int) []int {
	temp := []int{num}
	temp = append(temp, nums[:]...)
	copyTemp := make([]int, len(temp))
	for i := 0; i < len(temp); i++ {
		copyTemp[getPosL(i, len(temp))] = temp[i]
	}
	return copyTemp
}

//  向右移动两位的位置
func getPosL(pos, length int) int {
	return (pos + 2) % length
}
