package main

func pow(a, b int) int {
	return a*a + b*b
}
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	// 需要找到对角线
	temp, index := 0, 0
	m := make([]int, 5)
	m[0] = pow(p1[0]-p2[0], p1[1]-p2[1])
	m[1] = pow(p2[0]-p3[0], p2[1]-p3[1])
	m[2] = pow(p3[0]-p1[0], p3[1]-p1[1])
	// 此时的temp 是最大值
	for i := range m {
		if temp < m[i] {
			temp = m[i]
			index = i
		}
	}
	//fmt.Println(m)
	//  三种情况了
	if index == 0 { // p1,p2 为对角线
		if m[1]+m[2] != m[0] {
			return false
		}
		// 剩下的得重新算一手
		m[3] = pow(p2[0]-p4[0], p2[1]-p4[1])
		m[4] = pow(p4[0]-p1[0], p4[1]-p1[1])
		if m[3]+m[4] != m[0] {
			return false
		}

	} else if index == 1 { // p2, p3 为对角线
		if m[0]+m[2] != m[1] {
			return false
		}
		m[3] = pow(p2[0]-p4[0], p2[1]-p4[1])
		m[4] = pow(p4[0]-p3[0], p4[1]-p3[1])
		if m[3]+m[4] != m[1] {
			return false
		}
	} else if index == 2 { // p3, p1为对角线
		if m[0]+m[1] != m[2] {
			return false
		}
		m[3] = pow(p1[0]-p4[0], p1[1]-p4[1])
		m[4] = pow(p4[0]-p3[0], p4[1]-p3[1])
		if m[3]+m[4] != m[2] {
			return false
		}
	}

	// 瓜皮 4个点全是同一个
	if temp == 0 {
		return false
	}
	// 最后判断一手是否是正方形
	// 此时的temp 前一个值
	temp = m[4] // 3, 4 绝对不是最大值
	for i := range m {
		if i == index {
			continue
		}
		if temp != m[i] {
			return false
		}
		temp = m[i]
	}
	return true
}

func main() {
	p1 := []int{0, 0}
	p2 := []int{5, 0}
	p3 := []int{5, 4}
	p4 := []int{0, 4}
	validSquare(p1, p2, p3, p4)
}
