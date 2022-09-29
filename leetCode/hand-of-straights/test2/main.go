package main

import (
	"fmt"
	"sort"
)

//
//func isNStraightHand(hand []int, groupSize int) bool {
//	if len(hand)%groupSize != 0 {
//		return false
//	}
//	if groupSize == 1 {
//		return true
//	}
//	sort.Ints(hand)
//	//fmt.Println(hand)
//	visited := make([]bool, len(hand))
//	flag := false
//	var dfs func(index, num, now int)
//
//	dfs = func(index, num, now int) {
//		if index == len(hand) {
//			flag = true
//			return
//		}
//		if num == groupSize {
//			num = 0
//			now = 0
//			//fmt.Println(ans, visited)
//		}
//		for i := 0; i < len(hand); i++ {
//			if visited[i] {
//				continue
//			}
//			if now == 0 && num == 0 {
//				visited[i] = true
//				now = hand[i]
//				index++
//				num++
//			} else if hand[i] == now+1 {
//				visited[i] = true
//				now += 1
//				dfs(index+1, num+1, now)
//				now -= 1
//				visited[i] = false
//
//			}
//		}
//		// if index == len(hand) {
//		// 	flag = true
//		// 	return
//		// }
//		return
//
//	}
//
//	dfs(0, 0, 0)
//	//fmt.Println(ans)
//	return flag
//
//}
func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize > 0 {
		return false
	}
	cnts := map[int]int{}
	for _, h := range hand {
		cnts[h]++ // 打点计数
	}
	sort.Ints(hand)
	for _, h := range hand {
		// 进入这个循环，代表这个groupSize的第一个数了，不复合就无了呗
		for cnts[h] > 0 {
			//for i := h; i < h+groupSize; i++ {
			//	if cnts[i] == 0 {
			//		return false
			//	}
			//	cnts[i]--
			//}

			// 高配版：
			temp := cnts[h] // 减去最小牌的次数
			for i := h; i < h+groupSize; i++ {
				if cnts[i] < temp {
					return false
				}
				cnts[i] -= temp
			}
		}
	}
	return true
}
func main() {
	//fmt.Println(isNStraightHand([]int{1, 2, 3}, 1))
	////fmt.Println(isNStraightHand([]int{1, 2, 3, 4}, 1))
	////fmt.Println(isNStraightHand([]int{1, 2, 3, 4}, 2))
	////fmt.Println(isNStraightHand([]int{1, 3, 5, 8}, 2))

	//fmt.Println(isNStraightHand([]int{8, 10, 12}, 3))
	//fmt.Println(isNStraightHand([]int{9, 11, 13}, 3))
	//fmt.Println(isNStraightHand([]int{9, 10, 11}, 3))
	//fmt.Println(isNStraightHand([]int{1, 2, 3, 8, 5, 6, 7, 8}, 4))
	//fmt.Println(isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3))

	s := []int{5, 1, 0, 6, 4, 5, 3, 0, 8, 9}
	sort.Ints(s)
	fmt.Println(s)
	fmt.Println(isNStraightHand(s, 2))
}
