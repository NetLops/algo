package main

func numTeams(rating []int) int {
	smallLeft := 0
	smallRight := 0

	largeLeft := 0
	largeRight := 0
	n := len(rating)
	ans := 0
	for i := 1; i < n-1; i++ {
		smallLeft = 0
		for j := 0; j < i; j++ {
			if rating[j] < rating[i] {
				smallLeft += 1
			}
		}
		largeLeft = i - smallLeft

		smallRight = 0
		for j := i + 1; j < n; j++ {
			if rating[j] < rating[i] {
				smallRight += 1
			}
		}
		largeRight = n - 1 - i - smallRight

		ans += smallLeft*largeRight + smallRight*largeLeft
	}
	return ans
}
func main() {

}
