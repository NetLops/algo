package main

import (
	"fmt"
	"sort"
)

//func isNStraightHand(hand []int, groupSize int) bool {
//	//sort.Ints(hand)
//	flag := 0
//	m := map[int]int{}
//	for i := 0; i < len(hand); i++ {
//		// 转换成 例如 0,1,2,3,4 这种
//		m[hand[i]%groupSize]++
//		flag += hand[i] % 2
//	}
//	//fmt.Println(flag)
//	if flag == 0 || (len(hand) != 1 && flag == len(hand) && groupSize != 1) {
//		// 代表出现 全奇偶的情况
//		return false
//	}
//	keys := []int{}
//	prev := m[hand[0]%groupSize]
//	for k, v := range m {
//		// 如果v 的值不等
//		if prev != v {
//			return false
//		}
//		keys = append(keys, k)
//	}
//	// 如果 不等于 groupSize
//	if len(keys) != groupSize {
//		return false
//	}
//	sort.Ints(keys)
//	prekSum := keys[0] + keys[len(keys)-1]
//	for i := 0; i < len(keys)/2; i++ {
//		if prekSum != keys[i]+keys[len(keys)-i-1] {
//			return false
//		}
//	}
//	return true
//}

func isNStraightHand(hand []int, groupSize int) bool {
	// 贪心
	if len(hand)%groupSize != 0 {
		return false
	}
	sort.Ints(hand)
	m := map[int]int{}
	for i := 0; i < len(hand); i++ {
		m[hand[i]]++
	}
	for i := 0; i < len(hand); i++ {
		// 代表已经没了 0 0 1 1 2 2 // 执行到 0 0， 此时 1 1 2 2 都无了
		if m[hand[i]] == 0 {
			continue
		}
		for j := 1; j < groupSize; j++ {
			if m[hand[i+j]] == 0 {
				return false
			}
			m[hand[i+j]]--
		}
	}
	return true
}
func main() {
	fmt.Println(isNStraightHand([]int{1, 2, 3}, 1))
	//fmt.Println(isNStraightHand([]int{1, 2, 3, 4}, 1))
	//fmt.Println(isNStraightHand([]int{1, 2, 3, 4}, 2))
	//fmt.Println(isNStraightHand([]int{1, 3, 5, 8}, 2))
	fmt.Println(isNStraightHand([]int{8, 10, 12}, 3))
	fmt.Println(isNStraightHand([]int{9, 11, 13}, 3))
	fmt.Println(isNStraightHand([]int{9, 10, 11}, 3))
}
