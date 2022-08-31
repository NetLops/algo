package main

import "fmt"

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
func isMatch(a, b, x int) bool {
	if abs(a-x) <= abs(b-x) && a < b {
		return true
	}
	return false
}
func findClosestElements(arr []int, k int, x int) []int {

	// 先找x 是否在 arr 中 二分法

	if x < arr[0] { // 代表在最左边
		return arr[0:k]
	} else if x > arr[len(arr)-1] { // 代表在最右边
		return arr[len(arr)-k : len(arr)]
	}
	left, right := 0, len(arr)-1
	for left < right {
		m := left + (right-left)>>1
		if x <= arr[m] {
			right = m
		} else {
			left = m + 1
		}
	}
	//fmt.Println(left, right)

	if arr[left] != x { // 代表得向左移动一位
		if isMatch(arr[left-1], arr[left], x) {
			left--
		}
	}
	// 此时就需要双指针措施了
	// 根据先左后有的顺序，双指针一左一右的来
	left, right = left, left+1
	ans := make([]int, k)
	for k != 0 {
		if left >= 0 && right <= len(arr)-1 {
			if isMatch(arr[left], arr[right], x) {
				//arrLeft = append(arrLeft, arr[left])
				left--
			} else {
				//arrRight = append(arrRight, arr[right])
				right++
			}
		} else if left >= 0 {
			//arrLeft = append(arrLeft, arr[left])
			left--
		} else if right <= len(arr)-1 {
			//arrRight = append(arrRight, arr[right])
			right++
		}
		k--
	}
	for i := left + 1; i < right; i++ {
		ans[k] = arr[i]
		k++
	}
	return ans
}
func main() {
	fmt.Println(findClosestElements([]int{0, 1, 1, 1, 2, 3, 6, 7, 8, 9}, 9, 4))
}

//[0,1,1,1,2,3,6,7,8,9]
//9
//4
