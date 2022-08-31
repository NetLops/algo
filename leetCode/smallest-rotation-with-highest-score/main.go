package main

//
//func max(a, b int) int {
//	if a < b {
//		return b
//	}
//	return a
//}
//func bestRotation(nums []int) int {
//	//ans := 0
//	maxInt := 0
//	maxIndex := 0
//	for i := 0; i < len(nums); i++ {
//		temp := moveNums(i, nums)
//		if maxInt < getScore(temp) {
//			maxInt = getScore(temp)
//			maxIndex = i
//		}
//	}
//	return maxIndex
//}
//func moveNums(k int, nums []int) []int {
//	temp := make([]int, len(nums))
//	for i := 0; i < len(nums); i++ {
//		if i-k < 0 {
//			temp[len(nums)+i-k] = nums[i]
//		} else {
//			temp[i-k] = nums[i]
//		}
//	}
//
//	//j := k
//	//tail := 0
//	//for i := 0; i < len(temp); i++ {
//	//	if j <= len(temp)-1 {
//	//		temp[i] = nums[j]
//	//		j++
//	//	} else {
//	//		temp[i] = nums[tail]
//	//		tail++
//	//	}
//	//
//	//}
//	return temp
//}
//func getScore(temp []int) int {
//	score := 0
//	for i := 0; i < len(temp); i++ {
//		if temp[i] <= i {
//			score++
//		}
//	}
//	return score
//}

// ------------------------
//func add(l, r int, temp *[]int) {
//	(*temp)[l]++
//	(*temp)[r+1]--
//}
//func bestRotation(nums []int) int {
//	temp := make([]int, 100010)
//	n := len(nums)
//	for i := 0; i < n; i++ {
//		a := (i - (n - 1) + n) % n
//		b := (i - nums[i] + n) % n
//		if a <= b {
//			add(a, b, &temp)
//		} else {
//			add(0, b, &temp)
//			add(a, n-1, &temp)
//		}
//	}
//	for i := 1; i <= n; i++ {
//		temp[i] += temp[i-1]
//	}
//	ans := 0
//	k := temp[0]
//	for i := 1; i <= n; i++ {
//		if temp[i] > k {
//			k = temp[i]
//			ans = i
//		}
//	}
//	return ans
//}
// ------------------------
func add(left, right int, kArr *[]int) {
	(*kArr)[left]++
	(*kArr)[right+1]--

}
func bestRotation(nums []int) int {
	kArr := make([]int, len(nums)+1) // 那个差分 会 在 right + 1
	n := len(nums)
	for i := 0; i < n; i++ {
		x := nums[i]
		if i < x { // 一段 [i + 1,i -x + n]
			add(i+1, i-x+n, &kArr)
		} else { // i >= x 两段 [0,i -x], [i+1, n-1]
			add(0, i-x, &kArr)
			add(i+1, n-1, &kArr)
		}
	}
	maxK := 0
	max := 0
	cur := 0
	for i := 1; i < n; i++ {
		cur += kArr[i]
		if max < cur {
			maxK = i
			max = cur
		}
	}
	// 输出前缀和
	//for i := 1; i < len(kArr); i++ {
	//	kArr[i] += kArr[i-1]
	//}
	//fmt.Println(kArr)
	return maxK
}
func main() {
	//fmt.Println(bestRotation([]int{2, 3, 1, 4, 0}))
	bestRotation([]int{1, 3, 0, 2, 4})

}
