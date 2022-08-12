package main

import (
	"fmt"
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	dp := make([]string, len(nums))
	for i := range nums {
		dp[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(dp, func(i, j int) bool {
		if dp[i]+dp[j] < dp[j]+dp[i] {
			return true
		} else {
			return false
		}
	})
	temp := ""
	for i := range dp {
		temp += dp[i]
		fmt.Print(dp[i], " ")
	}
	return temp
}
func main() {
	fmt.Println("50" < "5")
	fmt.Println("50" > "5")
	fmt.Println("50" > "51")
	fmt.Println("501" > "51")
	fmt.Println("501" > "511")
	fmt.Println("501" > "50")
	fmt.Println("50"+"5" > "5"+"50")
	fmt.Println("50"+"5" < "5"+"50")
	fmt.Println("330"+"33" > "33"+"330")
	fmt.Println("330"+"33" < "33"+"330")
}
