package main

type pair struct {
	start int
	end   int
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// two pointer
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	ans := [][]int{}
	i, j := 0, 0
	for i < len(firstList) && j < len(secondList) { // 任意一个完了，就不用继续了
		maxStart := max(firstList[i][0], secondList[j][0])
		minEnd := min(firstList[i][1], secondList[j][1])
		if maxStart <= minEnd {
			ans = append(ans, []int{maxStart, minEnd})
		}
		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}
	return ans
}
func main() {

}
