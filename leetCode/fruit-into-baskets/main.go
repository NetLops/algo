package main

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//func totalFruit(fruits []int) int {
//	ans := 0
//	counts := make([]int, len(fruits)+20)
//	for i, j, t := 0, 0, 0; i < len(fruits); i++ {
//		counts[fruits[i]]++
//		if counts[fruits[i]] == 1 {
//			t++
//		}
//		for t > 2 {
//			counts[fruits[j]]--
//			if counts[fruits[j]] == 0 {
//				t--
//			}
//			j++
//		}
//		ans = Max(ans, i-j+1)
//	}
//	return ans
//}

//func totalFruit(fruits []int) int {
//	ans := 0
//	cnts := make([]int, len(fruits)) //  统计 fruits 种类
//	n := len(fruits)
//	// j 左边界
//	for i, j, t := 0, 0, 0; i < n; i++ {
//		cnts[fruits[i]]++
//		// 为 1的时候代表该种类是第一次加入
//		if cnts[fruits[i]] == 1 {
//			t++
//		}
//		for t > 2 { // 超过两个种类
//			// j 左边界向右移动
//			// 减去当前类型的值
//			cnts[fruits[j]]--
//			if cnts[fruits[j]] == 0 { // 等于0就代表当前类型跑完了
//				t--
//			}
//			j++
//		}
//		ans = Max(ans, i-j+1)
//	}
//	return ans
//}

func totalFruit(fruits []int) int {
	ans, n := 0, len(fruits)
	cnts := make([]int, n+1) // 计数器，题目 的种类 不会超过n的长度
	// 滑动窗口，i 一直前进，当t的种类大于2 的时候 ，j会向右边移动，此时j为左边界
	for i, j, t := 0, 0, 0; i < n; i++ {
		cnts[fruits[i]]++ // 增加种类的数量
		if cnts[fruits[i]] == 1 {
			t++ // 当前 新增一个种类
		}
		// 当t 大于2 的时候 就代表超出范围了
		for t > 2 {
			// j 向有移动
			cnts[fruits[j]]--
			// 种类为0 代表此时 的种类的数量减少了
			if cnts[fruits[j]] == 0 {
				t--
			}
			j++ // 右移动
		}
		ans = Max(ans, i-j+1)
	}
	return ans
}
